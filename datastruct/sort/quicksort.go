package main

import "fmt"

func quickSort(a []int, low, high int){
	if low < high{ // 防止陷入死循环
		m := partition(a, low, high)
		fmt.Printf("m: %d\n", m)
		quickSort(a, low, m - 1)
		quickSort(a, m + 1, high)
	}
}

func partition(a []int, low, high int ) int{
	pivot := a[low]
	for ; low < high; {
		// 从后向前
		for ; low < high && a[high] >= pivot; {
			high--
		}
		a[low] = a[high]
		//fmt.Printf("high:%d\n", high)

		// 从前向后
		for ; low < high && a[low] <= pivot; {
			low++
		}
		a[high] = a[low]
		//fmt.Printf("low:%d\n", low)
	}
	a[low] = pivot

	fmt.Printf("&a: %p, a: %v\n", a, a)
	return low
}

// 另一种方法
func quickSort2(a []int, left, right int){
	if left < right{
		m := partition2(a, left, right)
		quickSort2(a, left, m - 1)
		quickSort2(a, m + 1, right)
	}
}
func partition2(a []int, left, right int) int{
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++{
		if a[i] < a[pivot]{
			swap(a, i, index)
			index++
		}
	}
	swap(a, pivot, index - 1) // a[index - 1] 是 <= a[pivot]    a[index...] > a[pivot]
	return index - 1
}
func swap(a[]int, i, j int){
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func main() {
	for {
		var K int
		fmt.Println("--------")
		fmt.Scanln(&K)
		a := []int{2, 1, 3, 4, 0}
		findKth(a, 5, K)
	}

	//quickSort(a, 0, 4)
	//quickSort(a, 0, len(a) - 1)
	//fmt.Println(a)
}
