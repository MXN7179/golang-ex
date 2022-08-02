package main

import "fmt"

func main() {
	var c byte
	count := 0
	for {
		//fmt.Println("c:")
		fmt.Scanf("%c", &c)

		if c == 32{
			count = 0
			continue
		}
		if c == '\r'{
			break
		}
		count++
		//fmt.Printf("c: %c, count: %d\n", c, count)
	}
	fmt.Println(count)
}

