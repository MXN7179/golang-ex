package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	x := 0
	y := 0
	var str string
	fmt.Scanf("%s\n", &str)
	subs := strings.Split(str, ";")
	l := len(subs)
	for i := 0; i < l; i++{
		//fmt.Println(subs[i])
		res, b := isValid(subs[i])
		if b == true{
			switch subs[i][0] {
			case 'A':
				x -= res
			case 'D':
				x += res
			case 'W':
				y += res
			case 'S':
				y -= res
			}
		}
		//fmt.Println(x, y)
	}
	fmt.Printf("%d,%d\n", x, y)
}

func isValid(str string) (int, bool ) {
	l := len(str)
	if !(l == 2 || l == 3){  // l == 2 || l == 3
		return -1, false
	}
	if str[0] != 'A' && str[0] != 'D' && str[0]	!= 'W' && str[0] != 'S'{
		return -1, false
	}

	res := 0
	for i := 1; i < l; i++{
		if unicode.IsNumber(rune(str[i])) == false{
			return -1, false
		}
	}
	res, _ = strconv.Atoi(str[1:])

	return res, true
}