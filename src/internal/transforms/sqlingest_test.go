package transforms

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/pachyderm/pachyderm/v2/src/internal/dockertestenv"
	"github.com/pachyderm/pachyderm/v2/src/internal/randutil"
	"github.com/pachyderm/pachyderm/v2/src/internal/testutil"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestSQLIngest(t *testing.T) {
	ctx := context.Background()
	inputDir, outputDir := t.TempDir(), t.TempDir()
	u := dockertestenv.NewMySQLURL(t)

	// load  DB
	db := testutil.OpenDBURL(t, u, dockertestenv.MySQLPassword)
	_, err := db.Exec(`CREATE TABLE test_data (
			id SERIAL PRIMARY KEY,
			col_a VARCHAR(100)
	)`)
	require.NoError(t, err)
	const N = 100
	for i := 0; i < N; i++ {
		_, err := db.Exec(`INSERT INTO test_data (col_a) VALUES (?)`, randutil.UniqueString(""))
		require.NoError(t, err)
	}

	// write queries
	const Shards = 2
	for i := 0; i < Shards; i++ {
		name := fmt.Sprintf("%04d", i)
		// query would normally be different per shard
		query := "select * from test_data"
		err := ioutil.WriteFile(filepath.Join(inputDir, name), []byte(query), 0755)
		require.NoError(t, err)
	}

	err = SQLIngest(ctx, SQLIngestParams{
		Logger: logrus.StandardLogger(),

		InputDir:  inputDir,
		OutputDir: outputDir,

		URL:      u,
		Password: dockertestenv.MySQLPassword,
		Format:   "json",
	})
	require.NoError(t, err)

	// check the file exists
	dirEnts, err := os.ReadDir(outputDir)
	require.NoError(t, err)
	require.Len(t, dirEnts, Shards)
	const outputName = "0000"
	require.Equal(t, outputName, dirEnts[0].Name())
	lineCount := countLinesInFile(t, filepath.Join(outputDir, outputName))
	require.Equal(t, N, lineCount)
}

func countLinesInFile(t testing.TB, p string) int {
	f, err := os.Open(p)
	require.NoError(t, err)
	defer f.Close()
	br := bufio.NewReader(f)
	var count int
	for {
		ru, _, err := br.ReadRune()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		if ru == '\n' {
			count++
		}
	}
	return count
}