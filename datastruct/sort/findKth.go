package main

import "fmt"
//有一个整数数组，请你根据快速排序的思路，找出数组中第 k 大的数。
//给定一个整数数组 a ,同时给定它的大小n和要找的 k ，请返回第 k 大的数(包括重复的元素，不用去重)，保证答案存在。
//要求：时间复杂度 O(nlogn)，空间复杂度 O(1)
//数据范围：0 <= n ≤1000， 1≤K≤n，数组中每个元素满足 0≤val≤10000000

func main() {
	for {
		a := []int{2, 1, 3, 4, 0}
		n := 5
		var k int
		fmt.Println("------")
		fmt.Scanln(&k)
		findKth(a, n, k)
	}
}
func findKth( a []int ,  n int ,  K int ) int {

	res := findK(a, 0, n - 1, K - 1)
	fmt.Println("the Kth is", res)
	return 0
}

func findK(a []int, low, high, K int) int{
	var res = -1
	if low < high{ // 防止陷入死循环
		m := partitionGreater(a, low, high)
		fmt.Printf("m: %d\n", m)
		if m == K{
			fmt.Println("m == k")
			res = a[m]
			return res
		}else if m > K{
			fmt.Println("m > k")
			if low == m - 1{
				res = a[low]
				return res
			}
			res = findK(a, low, m - 1, K)
		}else{
			fmt.Println("m < k")
			if high == m + 1{
				res = a[high]
				return res
			}
			res = findK(a, m + 1, high, K)
		}
	}
	return res
}

func partitionGreater(a []int, low, high int ) int{
	pivot := a[low]
	for ; low < high; {
		// 从后向前
		for ; low < high && a[high] <= pivot; {
			high--
		}
		a[low] = a[high]
		//fmt.Printf("high:%d\n", high)

		// 从前向后
		for ; low < high && a[low] >= pivot; {
			low++
		}
		a[high] = a[low]
		//fmt.Printf("low:%d\n", low)
	}
	a[low] = pivot

	fmt.Printf("&a: %p, a: %v\n", a, a)
	return low
}
