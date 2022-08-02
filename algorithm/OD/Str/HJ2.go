package main

import "fmt"
func toLittle(c byte) (r byte){
	r = c
	if r >= 'A' && r <= 'Z'{
		r += 32
	}
	return
}
func main() {
	var str string
	var c byte
	count := 0
	for{
		fmt.Scanf("%c", &c)
		if c == '\r'{
			break
		}
		r := toLittle(c)
		str += string(r)
	}
	fmt.Scanf("\n%c", &c)
	c = toLittle(c)

	//fmt.Println(str, c)
	lens := len(str)
	for i := 0; i < lens; i++{
		if  str[i] == c {
				count++
		}
	}
	fmt.Println(count)
}