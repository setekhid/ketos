package main

import (
	"github.com/setekhid/ketos/pkg/rootpath"
)

import "C"

func expandPathName(pathNameC *C.char, ro bool) *C.char {

	pathNameG := C.GoString(pathNameC)
	pathNameG = rootpath.ExpandPath(pathNameG, ro)
	pathNameC = C.CString(pathNameG)

	return pathNameC
}

func main() {}
