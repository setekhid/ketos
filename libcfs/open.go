package main

import (
	"github.com/rainycape/dl"
	"log"
)

import "C"

//export open
func open(path *C.char, flags C.int, mode *C.int) *C.int {

	path = expandPathName(path)

	libc, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer libc.Close()

	var libc_open func(*C.char, C.int, *C.int) *C.int
	err = libc.Sym("open", &libc_open)
	if err != nil {
		log.Fatalln(err)
	}

	return libc_open(path, flags, mode)
}
