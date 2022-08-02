package main

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX = 10000 + 1
//var h [MAX]int
func _min(a, b int) (m int){
	m = a
	if a > b{
		m = b
	}
	return
}
func toDigit(d []string) []int{
	l := len(d)
	r := make([]int, l)
	for i := 0; i < l; i++{
		r[i], _ = strconv.Atoi(d[i])
	}
	return r
}
func main() {
	var str string
	fmt.Scanf("%s", &str)
	h := strings.Split(str, ",")
	r := toDigit(h)
	//fmt.Println(r)
	l := len(r)

	max := 0
	for i := 0; i < l; i++{
		for j := i + 1; j < l; j++{
			min := _min(r[i], r[j])
			cur := min * (j - i)
			if max < cur{
				max = cur
			}
		}
	}
	fmt.Println(max)
}
