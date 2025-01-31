package pfsutil

import (
	"github.com/pachyderm/pachyderm/v2/src/internal/client"
	"github.com/pachyderm/pachyderm/v2/src/pfs"
)

func MetaCommit(commit *pfs.Commit) *pfs.Commit {
	return client.NewSystemRepo(commit.Repo.Project.GetName(), commit.Repo.Name, pfs.MetaRepoType).NewCommit(commit.Branch.Name, commit.Id)
}
