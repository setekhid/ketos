package rootpath_test

import (
	"github.com/setekhid/ketos/pkg/rootpath"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestOverlayFS(t *testing.T) {

	rootfs, err := rootpath.NewOverlayFS(
		"./testdata/lowers/a",
		"./testdata/lowers/b",
		"./testdata/lowers/c",
		"./testdata/top",
	)
	require.NoError(t, err)

	asset1, err := rootfs.Expand("/asset1")
	require.NoError(t, err)
	t.Log(asset1)
	assert.True(t, strings.HasSuffix(asset1, "testdata/lowers/b/asset1"))
	assert.True(t, strings.HasPrefix(asset1, "/"))
}
