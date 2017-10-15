package main

import "C"
import "fmt"

//export say_hello
func say_hello(name *C.char, len C.int) {
	go_name := C.GoStringN(name, len)
	fmt.Printf("nautilus says: hello, %s!\n", go_name)
}

//export say_bye
func say_bye() {
	fmt.Print("nautilus says: bye!\n")
}

func main() {
}
