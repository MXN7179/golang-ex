package main

import (
	"fmt"
	"math"
)

// 记忆化搜索
// n个物品 (0-i)  (w,v)
//0: (2, 3)
//1: (1, 2)
//2: (3, 4)
//3: (2, 2)

const MAX = 100

var (
	W  uint = 5
	n  uint = 4
	w       = [MAX]uint{2, 1, 3, 2}
	v       = [MAX]uint{3, 2, 4, 2}
	dp [MAX][MAX]uint
)

func rec(i, j uint) uint {
	var res uint
	if i == n {
		res = 0
	} else if j < w[i] {
		res = rec(i+1, j)
	} else {
		res = uint(math.Max(float64(rec(i+1, j)), float64(rec(i+1, j-w[i])+v[i])))
	}
	return res
}

// 记忆化搜索
// n个物品 (0-i)  (w,v)
func main() {
	for j := uint(0); j <= W; j++ { // 初始化 dp[n][j]
		dp[n][j] = 0
	}

	res := rec(0, W)
	fmt.Println(res)
}
