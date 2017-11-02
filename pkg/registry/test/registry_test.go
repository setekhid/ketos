package test

import (
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestRegistryBasicUsage(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	// two image repository, local and remote
	image := registry.DockerImage("alpine:3.6")
	image0 := registry.DockerImage("localhost:5000/tasting/alpine:3.6")

	// connecting registries
	repo, tag, err := image.Connect()
	require.NoError(t, err)
	repo0, tag0, err := image0.Connect()
	require.NoError(t, err)

	assert.Equal(t, tag, tag0)
	assert.Equal(t, tag, "3.6")

	// get manifest from remote
	manifest, err := repo.GetManifest(tag)
	require.NoError(t, err)

	// transmit layers
	for _, layer := range manifest.FSLayers {

		func() {

			pipeR, pipeW, err := os.Pipe()
			require.NoError(t, err)
			defer pipeR.Close()

			go func() {
				defer pipeW.Close()

				err = repo.GetLayer(layer.BlobSum, pipeW)
				require.NoError(t, err)
			}()

			err = repo0.PutLayer(layer.BlobSum, pipeR)
			require.NoError(t, err)
		}()
	}

	// put manifest to local
	err = repo0.PutManifest(tag, manifest)
	require.NoError(t, err)
}
