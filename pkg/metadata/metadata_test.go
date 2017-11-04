package metadata_test

import (
	"io/ioutil"
	"os"
	"testing"

	manifestV1 "github.com/docker/distribution/manifest/schema1"
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetadataBasicUsage(t *testing.T) {

	tempDir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(tempDir)

	meta, err := metadata.NewMetadata(tempDir, "setekhid/dummy:latest")
	require.NoError(t, err)

	// check repository config
	conf, err := meta.GetConfig()
	require.NoError(t, err)
	assert.Equal(t, "setekhid/dummy:latest", conf.InitImageName)

	// check manifest storage
	manifest := &manifestV1.Manifest{
		Name: "setekhid/dummy",
		Tag:  "latest",
	}
	err = meta.PutManifest("latest", manifest)
	assert.NoError(t, err)
	manifest0, err := meta.GetManifest("latest")
	assert.Equal(t, manifest, manifest0)

	// check tags listing
	tags, err := meta.ListTags()
	require.NoError(t, err)
	assert.EqualValues(t, []string{"latest"}, tags)
}
