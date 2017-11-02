package registry_test

import (
	"github.com/setekhid/ketos/pkg/registry"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerImage(t *testing.T) {

	type ImagePair struct {
		Registry   string
		Repository string
		Tag        string
	}
	hubUrl := registry.DefaultRegistry

	cases := map[string]ImagePair{
		"setekhid/ketos":             {hubUrl, "setekhid/ketos", "latest"},
		"quay.io/setekhid/ketos:0.1": {"quay.io", "setekhid/ketos", "0.1"},
		"alpine:3.6":                 {hubUrl, "library/alpine", "3.6"},
	}

	for name, pair := range cases {

		registry, repository, tag := registry.DockerImage(name).Split()
		assert.EqualValues(t, pair, ImagePair{registry, repository, tag})
	}
}
