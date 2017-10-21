package main

import (
	"github.com/rainycape/dl"
	"log"
)

import "C"

//export eaccess
func eaccess(path *C.char, mode C.int) *C.int {

	path = expandPathName(path)

	libc, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer libc.Close()

	var libc_eaccess func(*C.char, C.int) *C.int
	err = libc.Sym("eaccess", &libc_eaccess)
	if err != nil {
		log.Fatalln(err)
	}

	return libc_eaccess(path, mode)
}
