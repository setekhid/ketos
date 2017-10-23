package registry

import hreg "github.com/heroku/docker-registry-client/registry"

var (
	defaultRegistry = "https://registry-1.docker.io/"
	defaultName     = ""
	defaultPasswd   = ""
)

func NewRegitry(registry, name, passwd string) (*hreg.Registry, error) {
	if registry == "" {
		registry = defaultRegistry
	}

	if name == "" {
		name = defaultName
	}

	if passwd == "" {
		passwd = defaultPasswd
	}

	hub, err := hreg.New(registry, name, passwd)
	if err != nil {
		return nil, err
	}

	return hub, err
}
