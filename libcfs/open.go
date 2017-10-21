package main

import (
	"github.com/rainycape/dl"
	"log"
	"os"
)

import "C"

//export open
func open(path *C.char, flags C.uint, mode C.int) *C.int {

	ro := flags&C.uint(os.O_WRONLY|os.O_RDWR|os.O_APPEND|os.O_CREATE) == 0
	path = expandPathName(path, ro)

	libc, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer libc.Close()

	var libc_open func(*C.char, C.uint, C.int) *C.int
	err = libc.Sym("open", &libc_open)
	if err != nil {
		log.Fatalln(err)
	}

	return libc_open(path, flags, mode)
}
