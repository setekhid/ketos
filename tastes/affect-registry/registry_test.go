package registry_test

import (
	"github.com/heroku/docker-registry-client/registry"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	url = "http://registry-1.docker.io/"
	username = ""
	password = ""
)

func TestRepositoryTags(t *testing.T) {
	t.SkipNow()

	hub, err := registry.New(url, username, password)
	assert.NoError(t, err)

	tags, err := hub.Tags("library/docker")
	assert.NoError(t, err)

	assert.Contains(t, tags, "stable-dind")
	assert.Contains(t, tags, "17.05-dind")
}

func TestManifests(t *testing.T) {

	hub, err := registry.New(url, username, password)
	assert.NoError(t, err)

	manifest, err := hub.Manifest("library/docker", "17.05-dind")
	assert.NoError(t, err)
	json, err := manifest.MarshalJSON()
	assert.NoError(t, err)
	t.Log(string(json))
}
