package main

import "fmt"

func df1() {

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	i := 1
	defer func() {
		i++
		fmt.Println("defer:", i)
	}()
	i++
	fmt.Println(i)
	return
}
func main() {
	df1()
}
