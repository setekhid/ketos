package pull

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testImageName = "library/alpine"
	testImageTag  = "latest"
)

func TestPull(t *testing.T) {
	assert := assert.New(t)

	err := pull(testImageName, testImageTag)
	assert.NoError(err)
}
