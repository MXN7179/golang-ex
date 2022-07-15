package main

import "fmt"

func _float() {
	f1 := 1e15
	f2 := 1e16

	fmt.Println(f1 == f2)
}

func main() {
	_float()
}
