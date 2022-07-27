package main

import "fmt"

var res = make([]rune, 0)

func Perm(a []rune, f func([]rune)) {
	perm(a, f, 0)
}
// 对索引 i 从 0 到 len(a) - 1，实现递归函数 perm().
func perm(a []rune, f func([]rune), i int) {
	// TODO
	index := i
	if index == len(a){
		//fmt.Println("end:", res)
		f(a)
		return
	}
	for i := index; i < len(a); i++ {
		// swap
		swap(a, index, i) //
		res = append(res, a[index])
		//fmt.Println("after append", res)

		// step
		perm(a, f, index+1)

		// reswap
		swap(a, index, i)

		// remove
		res = res[:len(res) - 1]
		//fmt.Println("after remove", res)
	}
}
func swap(s []rune, i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}
func main() {
	Perm([]rune("ABC"), func(a []rune) {
		fmt.Println(string(a))
	})
}