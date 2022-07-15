package main

import "fmt"

var tmp [100]int

func fib(n uint) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if tmp[n] != 0 {
		return tmp[n]
	}
	tmp[n] = fib(n-1) + fib(n-2)
	return tmp[n]
}

func main() {
	var n uint
	fmt.Println("n:")
	fmt.Scanln(&n)
	res := fib(n)
	fmt.Println(res)
}
