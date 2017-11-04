package image

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowImages(t *testing.T) {

	if testing.Short() {
		t.SkipNow()
	}

	assert := assert.New(t)

	err := showImages(defaultDir)
	assert.NoError(err)
}
