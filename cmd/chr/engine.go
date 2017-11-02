package main

import (
	"github.com/pkg/errors"
	"os"
)

type ChrootExecutor interface {
	Execute(repoPath, tagName string, userCommand []string) error
}

type ExecutorFunc func(repoPath, tagName string, userCommand []string) error

func (f ExecutorFunc) Execute(repoPath, tagName string,
	userCommand []string) error {

	return f(repoPath, tagName, userCommand)
}

func NewChrootExecutor(engineName string) (ChrootExecutor, error) {

	factory, exists := executorFactories[engineName]
	if !exists {
		return nil, errors.New("engine factory doesn't exists")
	}

	return factory(os.Environ())
}

type ChrootExecutorFactory func(env []string) (ChrootExecutor, error)

var (
	executorFactories = map[string]ChrootExecutorFactory{}
)

func AddExecutor(name string, factory ChrootExecutorFactory) {
	executorFactories[name] = factory
}
