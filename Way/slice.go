package main

import (
	"fmt"
	"sort"
)

func f1() {
	var s = [5]string{"a", "b", "cgo", "d"}
	for i := range s {
		fmt.Println(i, ":", s[i])
	}
}

func f2() {
	var s = new([6]string)
	s[0] = "1"
	s[5] = "2"
	fmt.Println(s)
}

func p1(s *[3]int) {
	fmt.Println(s)
}
func f3() {
	for i := 0; i < 3; i++ {
		p1(&[3]int{i, i * i, i * i * i})
	}
}

func f4() {
	s1 := make([]int, 10, 20)
	s2 := new([20]int)[0:10]

	fmt.Println(s1, s2)

}

func f5() {
	s := []int{1, 2, 3}
	s1 := s[1:1]
	s2 := s[1:2]
	fmt.Println(len(s1), ",", len(s2))
}

func RemoveFromSlice(s []int, start int, end int) []int { // end - start + 1 个
	if s == nil {
		return nil
	}
	rmLen := end - start + 1
	res := make([]int, len(s)-rmLen)
	at := copy(res, s[:start])
	copy(res[at:], s[end+1:])
	return res
}

func appRemove() {
	s := []int{1, 2, 3, 4, 5}
	res := RemoveFromSlice(s, 3, 4)
	fmt.Println(res, len(res))
}

func filterSlice(s []int, fn func(int) bool) []int {
	if s == nil {
		return nil
	}
	res := make([]int, 0)
	for i := range s {
		if fn(s[i]) == true {
			res = append(res, s[i])
		}
	}
	return res
}
func even(a int) bool {
	if a%2 == 0 {
		return true
	}
	return false
}
func appFilter() {
	s := []int{1, 2, 3, 4, 5}
	res := filterSlice(s, even)
	fmt.Println(res)
}

func sliceString() {
	s := "abcdef"
	si := []rune(s)

	fmt.Println(si)

	str1 := "LNMO"
	s1 := []byte(str1)
	fmt.Println(s1)

}

func SortSlice(s []float64, v float64) {
	sort.Float64s(s)
	sort.SearchFloat64s(s, v)

}

func splitString(str string, i int) (s1, s2 string) {
	return str[:i], str[i:]
}
func appSplitString() {
	fmt.Println(splitString("google", 2))
}

func reverseString(str string) string {
	s := []byte(str)
	i := 0
	j := len(s) - 1
	for i < j {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp

		i++
		j--
	}
	news := string(s)
	return news
}

func appReverseString() {
	str := "paner"
	fmt.Println(reverseString(str))
}

func uniqueString(str string) string {
	return ""
}

func swap(s []byte, i int) {
	tmp := s[i]
	s[i] = s[i+1]
	s[i+1] = tmp
}
func bubbleString(str string) string {
	s := []byte(str)
	n := len(s)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if s[j] > s[j+1] {
				swap(s, j)
			}
		}
	}
	return string(s)
}
func appBubble() {
	fmt.Println(bubbleString(""))
}

func d10(i int) int {
	return 10 * i
}
func mapFunction(fn func(i int) int, s []int) []int {
	res := make([]int, len(s))
	for i := 0; i < len(res); i++ {
		res[i] = fn(s[i])
	}
	return res
}

func main() {
	//f5()
	//appRemove()
	//appFilter()
	//sliceString()
	//appSplitString()
	//appReverseString()
	appBubble()
}
