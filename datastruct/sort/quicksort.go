package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth( a []int ,  n int ,  K int ) int {
	// write code here
	quickSort(a, 0, len(a) - 1)
	fmt.Println(a)
	return 0
}


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
		fmt.Printf("high:%d\n", high)

		// 从前向后
		for ; low < high && a[low] <= pivot; {
			low++
		}
		a[high] = a[low]
		fmt.Printf("low:%d\n", low)

		fmt.Printf("&a: %p, a: %v\n", a, a)
	}
	a[low] = pivot

	return low
}


func main() {
	a := []int{2, 1, 3, 4, 0}
	//findKth([]int{2, 1, 3, 4, 0}, 0, 0)
	//quickSort(a, 0, 4)
	quickSort(a, 0, len(a) - 1)
	fmt.Println(a)
}
