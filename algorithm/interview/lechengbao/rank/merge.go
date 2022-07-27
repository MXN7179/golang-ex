package main

import "fmt"

func merge(a []int, tmp []int, low1, high1, low2, high2 int){
	fmt.Printf("low1:%d high1:%d low2:%d high2:%d\n", low1, high1, low2, high2)
	if low1 > high1 || low2 > high2{
		return
	}
	i := low1
	j := low2
	tmpIndex := i

	for ; i <= high1 && j <= high2;{

		if a[i] <= a[j]{
			tmp[tmpIndex] = a[i]
			i++
		}else{
			tmp[tmpIndex] = a[j]
			j++
		}
		//fmt.Printf("tmp %d: %d\n", tmpIndex, tmp[tmpIndex])
		tmpIndex++
	}

	for ; i <= high1; i++{
		tmp[tmpIndex] = a[i]
		tmpIndex++
	}
	for ; j <= high2; j++{
		tmp[tmpIndex] = a[j]
		tmpIndex++
	}

	// 诶呀我去，我没有重新赋值，那 a 不还是原来的顺序吗，上层的排序就是乱的啊，想什么呢
	// tmp: low1 -> tmpIndex - 1   to  a: low1 -> high2
	for i := low1; i <= high2; i++{
		a[i] = tmp[i]
	}
}

func mergeSort(a []int, tmp []int, low, high int) {
	if low >= high{
		return
	}
	m := (low + high) / 2
	mergeSort(a, tmp, low, m)
	mergeSort(a, tmp, m + 1, high)
	merge(a, tmp, low, m, m + 1, high)
	fmt.Println("merge:", tmp)
}

func main() {
	a := []int{5, 1, 3, 4, 2}
	tmp := make([]int, len(a))
	mergeSort(a, tmp, 0, len(a) - 1 )
	//fmt.Println(a, tmp)
}