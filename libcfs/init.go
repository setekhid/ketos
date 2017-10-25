package main

import (
	"github.com/setekhid/ketos/pkg/rootpath"
)

var (
	RootLayers *rootpath.OverlayFS

	// TODO write log to files
	StdoutFile string
	StderrFile string
)

func init() {

	var err error

	RootLayers, err = rootpath.NewOverlayFSFromEnv()
	if err != nil {
		RootLayers = rootpath.NewDefaultOverlayFS()
	}
}
