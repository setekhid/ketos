package main

import (
	"syscall"
)

// #include <errno.h>
// void set_errno(int no) {
//     errno = no;
// }
import "C"

// set error number
func setErrno(err error) {

	if no, ok := err.(syscall.Errno); ok {
		C.set_errno(C.int(no))
	}

	C.set_errno(C.int(syscall.EIO))
}

func main() {}
