package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var str string
	fmt.Scanf("%s\n", &str)

	//for i := 0; i < len(str); i++{
	//	if !(str[i] == '.' || unicode.IsNumber(rune(str[i])) ){
	//		fmt.Println("NO")
	//		return
	//	}
	//}

	subs := strings.Split(str, ".")
	l := len(subs)
	if l < 4 || l > 4{
		fmt.Println("NO")
		return
	}
	//for i := 0; i < l; i++{
	//	if subs[i] == "" || ( len(subs[i])> 1 && subs[i][0] == '0') {
	//		fmt.Println("NO")
	//		return
	//	}
	//}
	ip := make([]int, len(subs))
	for i := 0; i < len(subs); i++{
		if unicode.IsDigit([]rune(subs[i]))

		ip[i], _= strconv.Atoi(subs[i])
		if ip[i] > 255 || ip[i] < 0 {
			fmt.Println("NO")
			return
		}
		//if i == 0 && ip[i] < 1{
		//	fmt.Println("NO")
		//	return
		//}
	}
	fmt.Println("YES")
}
