package main

import (
	"fmt"
	"os"
)

const MAXN = 500

var (
	path [MAXN][MAXN]bool
	data [MAXN]bool
	n int
)
func isVaild() bool{
	for i := 0; i < n; i++{
		if data[i] == false{
			return false
		}
	}
	return true
}
func isLoop() bool{
	m := make(map[int]bool, MAXN)
	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			if path[i][j] == true{
				m[j] = true
			}
		}
	}
	if len(m) == n{
		return false
	}

	// 默认开启的
	for i := 0; i < n; i++{
		if _, ok := m[i]; ok != true{
			data[i] = true
		}
	}
	return true
}

func judge(i int) bool{
	count := 0
	// open
	return false
	for j := 0; j < n; j++{
		if path[i][j] == true{
			data[j] = true
			count++
		}

	}

	if count == 0{
		v := isVaild()
		if v == true{ // 可行
			return true
		}else{   	// 不可行
			for j := 0; j < n; j++{
				if path[i][j] == true{
					data[j] = false
				}
			}
			return false
		}
	}

	j := 0
	for ; j < n; j++{
		if path[i][j] == true{
			v := judge(j)
			if v == true{
				fmt.Println("yes")
				os.Exit(0)
			}
		}
	}

	if j == n{ // 越界
		// close
		for j := 0; j < n; j++{
			if path[i][j] == true{
				data[j] = false
			}
		}
	}

	return false
}

func main() {
	//fmt.Println(path[0])
	fmt.Scanf("%d\n", &n)
	var s, d int
	for {
		_, err := fmt.Scanf("%d %d\n", &s, &d)
		if err != nil{
			break
		}
		path[s][d] = true
	}

	if r := isLoop(); r == false{
		fmt.Println("no")
		return
	}

	for i := 0; i < n; i++{
		judge(i)
	}

	fmt.Println("yes")
}

