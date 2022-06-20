package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	name := C.CString("Gopher")
	fmt.Printf("%p\n", name)
	fmt.Printf("%T\n", name)
	defer C.free(unsafe.Pointer(name))

	year := C.int(2008)
	fmt.Printf("%v\n", year)
	fmt.Printf("%T\n", year)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))
	fmt.Printf("%T\n", ptr)
	size := C.greet(name, year, (*C.char)(ptr))
	fmt.Printf("greeter function return type: %T\n", size)
	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))
}
