package pull

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testImageName = "library/alpine"
	testImageTag  = "latest"
)

func TestPullV2(t *testing.T) {
	assert := assert.New(t)

	err := pullV2(testImageName, testImageTag)
	assert.NoError(err)
}

func TestPullV1(t *testing.T) {
	assert := assert.New(t)

	err := pullV1(testImageName, testImageTag)
	assert.NoError(err)
}
