package main

import (
	"github.com/rainycape/dl"
)

import "C"

//export eaccess
func eaccess(cPath *C.char, cMode C.int) C.int {

	path, err := expandPathName(cPath, true)
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

	var libc_eaccess func(string, int) C.int
	err = libc.Sym("eaccess", &libc_eaccess)
	if err != nil {
		setErrno(err)
		return -1
	}

	return libc_eaccess(path, int(cMode))
}
