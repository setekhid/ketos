package main

import (
	"github.com/rainycape/dl"
	"os"
)

import "C"

//export open
func open(cPath *C.char, cFlags C.uint, cMode C.int) C.int {

	path := C.GoString(cPath)
	ro := cFlags&C.uint(os.O_WRONLY|os.O_RDWR|os.O_APPEND|os.O_CREATE) == 0

	expand := RootLayers.ExpandPath
	if !ro {
		expand = RootLayers.CopyForWriting
	}

	expanded, err := expand(path)
	if err != nil {
		setErrno(err)
		return -1
	}

	libc, err := dl.Open("libc", 0)
	if err != nil {
		setErrno(err)
		return -1
	}
	defer libc.Close()

	var libc_open func(string, uint, int) C.int
	err = libc.Sym("open", &libc_open)
	if err != nil {
		setErrno(err)
		return -1
	}

	return libc_open(expanded, uint(cFlags), int(cMode))
}
