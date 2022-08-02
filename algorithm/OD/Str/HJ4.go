package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s\n", &str)
	l := len(str)
	var i int
	for i = 0; i < l; i += 8{
		if i + 7 < l{
			s := str[i:i+8]
			fmt.Println(s)
		}
	}
	if i > l{
		i -= 8
		s := str[i:]
		b := 8 - len(s)
		for j := 0; j < b; j++{
			s += strconv.Itoa(0)
		}
		fmt.Println(s)
	}
}
