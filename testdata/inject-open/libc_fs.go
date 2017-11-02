// Based on article
// https://blog.gopheracademy.com/advent-2015/libc-hooking-go-shared-libraries/
package main

import "C"

import (
	"fmt"
	"log"

	"github.com/rainycape/dl"
)

//export open
func open(filename *C.char, flag C.int, mode C.int) *C.int {

	lib, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer lib.Close()

	var old_open func(*C.char, C.int, C.int) *C.int
	err = lib.Sym("open", &old_open)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("into golang hook of open\n")
	return old_open(filename, flag, mode)
}

func main() {}
