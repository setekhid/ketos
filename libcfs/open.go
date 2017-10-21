package main

import (
	"github.com/rainycape/dl"
	"log"
)

import "C"

//export open
func open(path *C.char, flags C.int, mode C.int) C.int {

	path = expandPathName(path)

	libc, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer libc.Close()

	var libc_open func(string, int, ...interface{}) int32
	err = libc.Sym("open", &libc_open)
	if err != nil {
		log.Fatalln(err)
	}

	return C.int(libc_open(C.GoString(path), int(flags), mode))
}
