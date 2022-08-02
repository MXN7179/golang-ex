package main

import "fmt"

func main() {
	m := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'A': 10,
		'B': 11,
		'C': 12,
		'D': 13,
		'E': 14,
		'F': 15,
	}

	var str string
	res := 0
	fmt.Scanf("%s\n", &str)
	str = str[2:]
	l := len(str)
	for i := 0; i < l; i++{
		res = res * 16 + m[str[i]]
	}
	fmt.Println(res)
}
