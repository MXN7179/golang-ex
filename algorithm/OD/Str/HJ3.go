package main

import "fmt"

const MAX = 500 + 1
func main() {
	arr := make([]int, MAX)
	var n, v int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++{
		fmt.Scanf("%d\n", &v)
		//fmt.Println(i, v)
		arr[v] = 1
	}
	for i := 0; i < MAX; i++{
		if arr[i] == 1{
			fmt.Println(i)
		}
	}
}
