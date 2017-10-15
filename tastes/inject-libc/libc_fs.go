// Based on article
// https://blog.gopheracademy.com/advent-2015/libc-hooking-go-shared-libraries/
package main

import "C"

import (
	"fmt"
	"log"

	"github.com/rainycape/dl"
)

//export strrchr
func strrchr(s *C.char, c C.int) *C.char {

	lib, err := dl.Open("libc", 0)
	if err != nil {
		log.Fatalln(err)
	}
	defer lib.Close()

	var old_strrchr func(s *C.char, c C.int) *C.char
	err = lib.Sym("strrchr", &old_strrchr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("into golang hook of strrchr\n")
	return old_strrchr(s, c)
}

func main() {}
