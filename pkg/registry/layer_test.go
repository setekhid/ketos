package registry_test

import (
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestSortFiles4Layer(t *testing.T) {

	original := []string{
		"abc/bbc/a.txt",
		"abc/bbc",
		"ccd",
	}
	expected := []string{
		"abc/bbc",
		"abc/bbc/a.txt",
		"ccd",
	}

	sorted := registry.SortFiles4Layer(original)

	assert.EqualValues(t, expected, sorted)
}

func TestLayerTarUntar(t *testing.T) {

	pipeR, pipeW, err := os.Pipe()
	require.NoError(t, err)
	defer pipeR.Close()

	// tar layer
	func() {
		defer pipeW.Close()

		digest, err := registry.TarLayerDirectory(pipeW, "./testdata")
		require.NoError(t, err)
		t.Log("layer digest:", digest.Hex())
	}()

	layerDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(layerDir)

	err = registry.UntarLayerDirectory(pipeR, layerDir)
	require.NoError(t, err)

	err = filepath.Walk(
		layerDir,
		func(path string, info os.FileInfo, err error) error {

			require.NoError(t, err)

			relPath, err := filepath.Rel(layerDir, path)
			require.NoError(t, err)

			origPath := filepath.Join("./testdata", relPath)
			t.Log("checking original path:", origPath)

			_, err = os.Stat(origPath)
			assert.NoError(t, err)

			return nil
		},
	)
	require.NoError(t, err)
}
