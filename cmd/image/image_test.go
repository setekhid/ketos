package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowImages(t *testing.T) {
	assert := assert.New(t)

	err := showImages(defaultDir)
	assert.NoError(err)
}
