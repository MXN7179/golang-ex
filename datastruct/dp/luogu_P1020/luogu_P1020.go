package main

import (
	"fmt"
	"io"
)

const MAX = 2e5 + 5

var (
	dp	[MAX]int
	a	[MAX]int
)

func main() {
	var data int
	for {
		_, err := fmt.Scan(&data)
		if err == io.EOF{
			break
		}
	}

	fmt.Println(data)
}