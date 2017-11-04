package metadata_test

import (
	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestKetosFolder(t *testing.T) {

	path, err := metadata.SeekKetosFolder("./testdata/a/c")
	require.NoError(t, err)
	assert.True(t, strings.HasSuffix(path, "/testdata/a/.ketos"))

	path, err = metadata.SeekKetosFolder("./testdata/a/b")
	require.NoError(t, err)
	assert.True(t, strings.HasSuffix(path, "/testdata/a/b/.ketos"))

	err = os.Chdir("./testdata/a")
	require.NoError(t, err)
	path, err = metadata.KetosFolder()
	require.NoError(t, err)
	assert.True(t, strings.HasSuffix(path, "/testdata/a/.ketos"))
}
