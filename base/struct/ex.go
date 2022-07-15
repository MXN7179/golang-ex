package main

import (
	"fmt"
	"unsafe"
)

func f1() {
	fmt.Println(unsafe.Sizeof(struct {
		//a uint8
		//c uint32
		//b uint16
	}{}))
}

func main() {
	f1()
}
