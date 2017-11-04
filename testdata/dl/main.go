package main

import (
	"bytes"
	"fmt"
	"github.com/rainycape/dl"
)

func main() {

	libc, err := dl.Open("libc", 0)
	if err != nil {
		panic(err)
	}
	defer libc.Close()

	var snprintf func([]byte, uint, string, ...interface{}) int
	err = libc.Sym("snprintf", &snprintf)
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 200)
	snprintf(buf, uint(len(buf)), "hello %s!\n", "world")
	s := string(buf[:bytes.IndexByte(buf, 0)])

	fmt.Println(s)
}
