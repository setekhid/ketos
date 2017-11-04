package push

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/setekhid/ketos/pkg/metadata"
	"github.com/stretchr/testify/assert"
)

var (
	testImageName = "127.0.0.1:5000/alpine"
	testImageTag  = "latest"
)

func TestPush(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	assert := assert.New(t)

	ketosFolder, err := metadata.KetosFolder()
	if err != nil {
		fmt.Println(err)
		return
	}

	defaultTagDir = filepath.Join(ketosFolder, "tags")
	defaultlLayer = filepath.Join(ketosFolder, "layers")

	err = push(testImageName, testImageTag)
	assert.NoError(err)
}
