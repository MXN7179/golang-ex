package main

import (
	"fmt"
	"strconv"
)

// s string 以 0 开头
// 每个段第一个字符【0-2】， 不能以0开头

var res = make([]string, 0, 4)

func IPSlice(s string, pn, i int) {
	//fmt.Printf("pn: %d, s[:%d]: %s, res: %s\n", pn, i, s[:i], res)

	if pn == 3 {
		if i+2 >= len(s) {
			res = append(res, s[i:])
			fmt.Printf("res: %s\n", res)
			res = res[:len(res)-1]
			return
		} else if i+3 >= len(s) && s[i] <= '2' && s[i] >= '0' {
			res = append(res, s[i:])
			fmt.Printf("res: %s\n", res)
			res = res[:len(res)-1]
			return
		} else {
			//fmt.Printf("xxxx %v, s[%d]: %s\n", res, i, s[:i])
			return
		}
	}

	//  索引step： 1 2 3
	res = append(res, s[i:i+1])
	IPSlice(s, pn+1, i+1)
	res = res[:len(res)-1]

	res = append(res, s[i:i+2])
	IPSlice(s, pn+1, i+2)
	res = res[:len(res)-1]

	if s[i] <= '2' && s[i] >= '0' {
		res = append(res, s[i:i+3])
		IPSlice(s, pn+1, i+3)
		res = res[:len(res)-1]
	}

}

func main() {
	//s := "2252252252"
	//s := "0250250252"
	//s := "2292252252"
	s := "3292252252"
	//s := "1"
	//s := "22292252252292252252"
	IPSlice(s, 0, 0)

	//toInt()
}

func toInt() {
	s := "0123"
	fmt.Println(strconv.Atoi(s))
}
