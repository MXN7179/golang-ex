package main

import "fmt"

//
// 假设为从大大到小有序

func main() {
	a := []int{10, 9, 4}
	b := []int{5, 4, 3}
	k := 5
	findK(a, b, k)
}

func findK(a, b []int, k int) (int, bool) {
	if a == nil || b == nil {
		return -1, false
	}

	pos := 1
	lenA := len(a)
	lenB := len(b)
	i := 0
	j := 0
	for i < lenA && j < lenB {
		if pos == k {
			var res int
			if a[i] > b[j] {
				res = a[i]
			} else {
				res = b[j]
			}
			fmt.Println(pos, res)
			return 0, true
		}

		if a[i] > b[j] {
			i++
			pos++
		} else {
			j++
			pos++
		}

	}

	var balance []int
	if i >= lenA {
		balance = b[j:]
	} else {
		balance = a[i:]
	}

	for i := 0; i < len(balance); i++ {
		if pos == k {
			fmt.Println(pos, balance[i])
			return 0, true
		}
		pos++
	}

	// k out of range
	if i >= len(balance) {
		fmt.Println("k out")
		return -1, false
	}

	return 0, true
}
