package cgo

// #include <stdlib.h>
import "C"

func Random() int {
	//return int(C.random())
	return 0
}

func Seed(i int) {
	//C.srandom(C.uint(i))
}

func main() {
	Random()
	Seed(10)

}
