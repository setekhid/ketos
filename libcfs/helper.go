package main

import (
	"github.com/setekhid/ketos/pkg/rootpath"
	"syscall"
)

// #include <errno.h>
// void set_errno(int no) {
//     errno = no;
// }
import "C"

// expand path to fake rootfs
func expandPathName(pathNameC *C.char, ro bool) (string, error) {

	pathNameG := C.GoString(pathNameC)
	return rootpath.ExpandPath(pathNameG, ro)
}

// set error number
func setErrno(err error) {

	if no, ok := err.(syscall.Errno); ok {
		C.set_errno(C.int(no))
	}

	C.set_errno(C.int(syscall.EIO))
}

func main() {}
