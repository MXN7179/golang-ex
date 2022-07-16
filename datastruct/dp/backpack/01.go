package main

import (
	"fmt"
	"math"
)

// n个物品 (0-i)  (w,v)
//0: (2, 3)
//1: (1, 2)
//2: (3, 4)
//3: (2, 2)

const MAX = 100

var (
	W  int = 5
	n  int = 4
	w      = [MAX]int{2, 1, 3, 2}
	v      = [MAX]int{3, 2, 4, 2}
	dp [MAX][MAX]int
)

// 记忆化搜索
func rec(i, j int) int {
	var res int
	if i == n {
		res = 0
	} else if j < w[i] {
		res = rec(i+1, j)
	} else {
		res = int(math.Max(float64(rec(i+1, j)), float64(rec(i+1, j-w[i])+v[i])))
	}
	return res
}
func initDp() {
	for j := int(0); j <= W; j++ { // 初始化 dp[n][j]
		dp[n][j] = 0
	}
}

// 动态规划
func solve() {
	initDp()

	for i := n - 1; i >= 0; i-- {
		for j := 0; j <= W; j++ {
			if j < w[i] {
				dp[i][j] = dp[i+1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i+1][j]), float64(dp[i+1][j-w[i]]+v[i])))
			}
		}
	}

	fmt.Println(dp[0][W])
}

func main() {
	//initdp()
	//res := rec(0, W)
	//fmt.Println(res)

	solve()
}
