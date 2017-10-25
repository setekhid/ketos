package main

import (
	"github.com/setekhid/ketos/pkg/rootpath"
)

var (
	RootLayers *rootpath.OverlayFS
)

func init() {

	var err error

	RootLayers, err = rootpath.NewOverlayFSFromEnv()
	if err != nil {
		RootLayers = rootpath.NewDefaultOverlayFS()
	}
}
