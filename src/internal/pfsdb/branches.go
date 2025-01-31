package pfsdb

import (
	"context"
	"fmt"
	"strings"

	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/pachsql"
	"github.com/pachyderm/pachyderm/v2/src/pfs"
)

const (
	getBranchBaseQuery = `
		SELECT
			branch.id,
			branch.name,
			branch.created_at,
			branch.updated_at,
			repo.id as "repo.id",
			repo.name as "repo.name",
			repo.type as "repo.type",
			project.id as "repo.project.id",
			project.name as "repo.project.name",
			commit.commit_set_id as "head.commit_set_id",
			repo.name as "head.repo.name",
			repo.type as "head.repo.type",
			project.name as "head.repo.project.name"
		FROM pfs.branches branch
			JOIN pfs.repos repo ON branch.repo_id = repo.id
			JOIN core.projects project ON repo.project_id = project.id
			JOIN pfs.commits commit ON branch.head = commit.int_id
	`
	getBranchByIDQuery   = getBranchBaseQuery + ` WHERE branch.id = $1`
	getBranchByNameQuery = getBranchBaseQuery + ` WHERE project.name = $1 AND repo.name = $2 AND repo.type = $3 AND branch.name = $4`
)

// SliceDiff takes two slices and returns the elements in the first slice that are not in the second slice.
// TODO this can be moved to a more generic package.
func SliceDiff[K comparable, V any](a, b []V, key func(V) K) []V {
	m := make(map[K]bool)
	for _, item := range b {
		m[key(item)] = true
	}
	var result []V
	for _, item := range a {
		if !m[key(item)] {
			result = append(result, item)
		}
	}
	return result
}

// GetBranchInfo returns a *pfs.BranchInfo by id.
func GetBranchInfo(ctx context.Context, tx *pachsql.Tx, id BranchID) (*pfs.BranchInfo, error) {
	branch := &Branch{}
	if err := tx.GetContext(ctx, branch, getBranchByIDQuery, id); err != nil {
		return nil, errors.Wrap(err, "could not get branch")
	}
	return fetchBranchInfoByBranch(ctx, tx, branch)
}

// GetBranchInfoByName returns a *pfs.BranchInfo by name
func GetBranchInfoByName(ctx context.Context, tx *pachsql.Tx, project, repo, repoType, branch string) (*pfs.BranchInfo, error) {
	row := &Branch{}
	if err := tx.GetContext(ctx, row, getBranchByNameQuery, project, repo, repoType, branch); err != nil {
		return nil, errors.Wrap(err, "could not get branch")
	}
	return fetchBranchInfoByBranch(ctx, tx, row)
}

// GetBranchID returns the id of a branch given a set strings that uniquely identify a branch.
func GetBranchID(ctx context.Context, tx *pachsql.Tx, project, repo, repoType, branch string) (BranchID, error) {
	var id BranchID
	if err := tx.GetContext(ctx, &id, `
		SELECT branch.id
		FROM pfs.branches branch
			JOIN pfs.repos repo ON branch.repo_id = repo.id
			JOIN core.projects project ON repo.project_id = project.id
		WHERE project.name = $1 AND repo.name = $2 AND repo.type = $3 AND branch.name = $4
	`, project, repo, repoType, branch); err != nil {
		return 0, errors.Wrapf(err, "could not get id for branch %v", branch)
	}
	return id, nil
}

// UpsertBranch creates a branch if it does not exist, or updates the head if the branch already exists.
// If direct provenance is specified, it will be used to update the branch's provenance relationships.
func UpsertBranch(ctx context.Context, tx *pachsql.Tx, branchInfo *pfs.BranchInfo) (BranchID, error) {
	if branchInfo.Branch.Repo.Name == "" {
		return 0, errors.Errorf("repo name required")
	}
	if branchInfo.Branch.Repo.Type == "" {
		return 0, errors.Errorf("repo type required")
	}
	if branchInfo.Branch.Repo.Project.Name == "" {
		return 0, errors.Errorf("project name required")
	}
	if branchInfo.Head.Id == "" {
		return 0, errors.Errorf("head commit required")
	}
	var branchID BranchID
	// TODO stop matching on pfs.commits.commit_id, because that will eventually be deprecated.
	// Instead, construct the commit_id based on existing project, repo, and commit_set_id fields.
	if err := tx.QueryRowContext(ctx,
		`
		INSERT INTO pfs.branches(repo_id, name, head)
		VALUES (
			(SELECT repo.id FROM pfs.repos repo JOIN core.projects project ON repo.project_id = project.id WHERE project.name = $1 AND repo.name = $2 AND repo.type = $3),
			$4,
			(SELECT int_id FROM pfs.commits WHERE commit_id = $5)
		)
		ON CONFLICT (repo_id, name) DO UPDATE SET head = EXCLUDED.head
		RETURNING id
		`,
		branchInfo.Branch.Repo.Project.Name,
		branchInfo.Branch.Repo.Name,
		branchInfo.Branch.Repo.Type,
		branchInfo.Branch.Name,
		CommitKey(branchInfo.Head),
	).Scan(&branchID); err != nil {
		return 0, errors.Wrap(err, "could not create branch")
	}
	// Compare current direct provenance to new direct provenance.
	newDirectProv := branchInfo.DirectProvenance
	oldDirectProv, err := GetDirectBranchProvenance(ctx, tx, branchID)
	if err != nil {
		return 0, errors.Wrap(err, "could not get direct branch provenance")
	}
	// Add net new direct provenance relationships.
	toAdd := SliceDiff[string, *pfs.Branch](newDirectProv, oldDirectProv, func(branch *pfs.Branch) string { return branch.Key() })
	toAddIDs := make([]BranchID, len(toAdd))
	for i, branch := range toAdd {
		toAddIDs[i], err = GetBranchID(ctx, tx, branch.Repo.Project.Name, branch.Repo.Name, branch.Repo.Type, branch.Name)
		if err != nil {
			return 0, errors.Wrap(err, "could not get to_id")
		}
	}
	if err := CreateDirectBranchProvenanceBatch(ctx, tx, branchID, toAddIDs); err != nil {
		return 0, errors.Wrap(err, "could not create branch provenance")
	}
	// Remove old direct provenance relationships that are no longer needed.
	toRemove := SliceDiff[string, *pfs.Branch](oldDirectProv, newDirectProv, func(branch *pfs.Branch) string { return branch.Key() })
	toRemoveIDs := make([]BranchID, len(toRemove))
	for i, branch := range toRemove {
		toRemoveIDs[i], err = GetBranchID(ctx, tx, branch.Repo.Project.Name, branch.Repo.Name, branch.Repo.Type, branch.Name)
		if err != nil {
			return 0, errors.Wrap(err, "could not get to_id")
		}
	}
	if err := DeleteDirectBranchProvenanceBatch(ctx, tx, branchID, toRemoveIDs); err != nil {
		return 0, errors.Wrap(err, "could not delete branch provenance")
	}
	return branchID, nil
}

// GetDirectBranchProvenance returns the direct provenance of a branch, i.e. all branches that it directly depends on.
func GetDirectBranchProvenance(ctx context.Context, tx *pachsql.Tx, id BranchID) ([]*pfs.Branch, error) {
	var branches []Branch
	if err := tx.SelectContext(ctx, &branches, `
		SELECT
			branch.id,
			branch.name,
			repo.name as "repo.name",
			repo.type as "repo.type",
			project.name as "repo.project.name"
		FROM pfs.branch_provenance bp
		    JOIN pfs.branches branch ON bp.to_id = branch.id
			JOIN pfs.repos repo ON branch.repo_id = repo.id
			JOIN core.projects project ON repo.project_id = project.id
		WHERE bp.from_id = $1
	`, id); err != nil {
		return nil, errors.Wrap(err, "could not get direct branch provenance")
	}
	var branchPbs []*pfs.Branch
	for _, branch := range branches {
		branchPbs = append(branchPbs, branch.Pb())
	}
	return branchPbs, nil
}

// GetBranchProvenance returns the full provenance of a branch, i.e. all branches that it either directly or transitively depends on.
func GetBranchProvenance(ctx context.Context, tx *pachsql.Tx, id BranchID) ([]*pfs.Branch, error) {
	var branches []Branch
	if err := tx.SelectContext(ctx, &branches, `
		WITH RECURSIVE prov(from_id, to_id) AS (
		    SELECT from_id, to_id
		    FROM pfs.branch_provenance
		    WHERE from_id = $1
		  UNION ALL
		    SELECT bp.from_id, bp.to_id
		    FROM prov JOIN pfs.branch_provenance bp ON prov.to_id = bp.from_id
		)
		SELECT DISTINCT
		    branch.id,
			branch.name,
			repo.name as "repo.name",
			repo.type as "repo.type",
			project.name as "repo.project.name"
		FROM pfs.branches branch
		    JOIN prov ON branch.id = prov.to_id
			JOIN pfs.repos repo ON branch.repo_id = repo.id
		    JOIN core.projects project ON repo.project_id = project.id
		WHERE branch.id != $1
	`, id); err != nil {
		return nil, errors.Wrap(err, "could not get branch provenance")
	}
	var branchPbs []*pfs.Branch
	for _, branch := range branches {
		branchPbs = append(branchPbs, branch.Pb())
	}
	return branchPbs, nil
}

// GetBranchSubvenance returns the full subvenance of a branch, i.e. all branches that either directly or transitively depend on it.
func GetBranchSubvenance(ctx context.Context, tx *pachsql.Tx, id BranchID) ([]*pfs.Branch, error) {
	var branches []Branch
	if err := tx.SelectContext(ctx, &branches, `
		WITH RECURSIVE subv(from_id, to_id) AS (
		    SELECT from_id, to_id
		    FROM pfs.branch_provenance
		    WHERE to_id = $1
		  UNION ALL
		    SELECT bp.from_id, bp.to_id
		    FROM subv JOIN pfs.branch_provenance bp ON subv.from_id = bp.to_id
		)
		SELECT DISTINCT
		    branch.id,
			branch.name,
			repo.name as "repo.name",
			repo.type as "repo.type",
			project.name as "repo.project.name"
		FROM pfs.branches branch
		    JOIN subv ON branch.id = subv.from_id
			JOIN pfs.repos repo ON branch.repo_id = repo.id
		    JOIN core.projects project ON repo.project_id = project.id
		WHERE branch.id != $1
	`, id); err != nil {
		return nil, errors.Wrap(err, "could not get branch provenance")
	}
	var branchPbs []*pfs.Branch
	for _, branch := range branches {
		branchPbs = append(branchPbs, branch.Pb())
	}
	return branchPbs, nil
}

// CreateBranchProvenance creates a provenance relationship between two branches.
func CreateDirectBranchProvenance(ctx context.Context, tx *pachsql.Tx, from, to BranchID) error {
	return CreateDirectBranchProvenanceBatch(ctx, tx, from, []BranchID{to})
}

// CreateBranchProvenanceBatch creates provenance relationships between a branch and a set of other branches.
func CreateDirectBranchProvenanceBatch(ctx context.Context, tx *pachsql.Tx, from BranchID, tos []BranchID) error {
	if len(tos) == 0 {
		return nil
	}
	query := `
		INSERT INTO pfs.branch_provenance(from_id, to_id)
		VALUES %s
		ON CONFLICT DO NOTHING
	`
	values := make([]string, len(tos))
	for i, to := range tos {
		values[i] = fmt.Sprintf("(%d, %d)", from, to)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))
	if _, err := tx.ExecContext(ctx, query); err != nil {
		return errors.Wrap(err, "could not add branch provenance")
	}
	return nil
}

// DeleteBranchProvenance deletes a provenance relationship between two branches.
func DeleteDirectBranchProvenance(ctx context.Context, tx *pachsql.Tx, from, to BranchID) error {
	return DeleteDirectBranchProvenanceBatch(ctx, tx, from, []BranchID{to})
}

// DeleteBranchProvenanceBatch deletes provenance relationships between a branch and a set of other branches.
func DeleteDirectBranchProvenanceBatch(ctx context.Context, tx *pachsql.Tx, from BranchID, tos []BranchID) error {
	if len(tos) == 0 {
		return nil
	}
	query := `
	DELETE FROM pfs.branch_provenance
	WHERE from_id = $1 AND to_id IN (%s);
	`
	values := make([]string, len(tos))
	for i, to := range tos {
		values[i] = fmt.Sprintf("%d", to)
	}
	query = fmt.Sprintf(query, strings.Join(values, ","))
	_, err := tx.ExecContext(ctx, query, from)
	return errors.Wrap(err, "could not delete branch provenance")
}

func fetchBranchInfoByBranch(ctx context.Context, tx *pachsql.Tx, branch *Branch) (*pfs.BranchInfo, error) {
	branchInfo := &pfs.BranchInfo{Branch: branch.Pb(), Head: branch.Head.Pb()}
	var err error
	branchInfo.DirectProvenance, err = GetDirectBranchProvenance(ctx, tx, branch.ID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get direct branch provenance")
	}
	branchInfo.Provenance, err = GetBranchProvenance(ctx, tx, branch.ID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get full branch provenance")
	}
	branchInfo.Subvenance, err = GetBranchSubvenance(ctx, tx, branch.ID)
	if err != nil {
		return nil, errors.Wrap(err, "could not get full branch subvenance")
	}
	return branchInfo, nil
}
