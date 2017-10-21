package client

import (
	"fmt"

	hreg "github.com/heroku/docker-registry-client"
)

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

	hub, err := hreg.New(url, name, passwd)
	if err != nil {
		return nil, fmt.Errorf("fail to new hub %s", err)
	}

	return hub, err
}
