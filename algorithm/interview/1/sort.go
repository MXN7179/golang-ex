package main

import "fmt"

// 1 ( 2 3 4)
//     2 (34) (4 3)
//	   3 (2 4

var (
	res = make([]int, 0)
)

func sort(s []int, begin int) {

	if begin == len(s){
		fmt.Println("end:", res)
		return
	}
	for i := begin; i < len(s); i++ {
		// swap
		swap(s, begin, i) //
		res = append(res, s[begin])
		//fmt.Println("after append", res)

		// step
		sort(s, begin+1)

		// reswap
		swap(s, begin, i)

		// remove
		res = res[:len(res) - 1]
		//fmt.Println("after remove", res)
	}
}

func swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func main() {

	sort([]int{0, 1, 2, 3}, 0)
}
