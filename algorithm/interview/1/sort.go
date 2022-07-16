package main

import "fmt"

// 1 ( 2 3 4)
//     2 (34) (4 3)
//	   3 (2 4

func sort(s []int, begin int) {

	for i := begin; i < len(s); i++ {
		if begin == len(s)-1 {
			fmt.Println("----")
		}
		// swap
		swap(s, begin, i) //
		//fmt.Printf("%d ", s[begin])
		// append
		sort(s, begin+1)
		swap(s, begin, i)
		// remove
	}
}

func swap(s []int, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func main() {

	sort([]int{1, 2, 3, 4}, 0)
}
