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

const MAX = 10

var (
	W  = 5
	n  = 4
	w  = [MAX]int{2, 1, 3, 2}
	v  = [MAX]int{3, 2, 4, 2}
	w2 = [MAX]int{0, 2, 1, 3, 2}
	v2 = [MAX]int{0, 3, 2, 4, 2}
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
func initDp(row int) {
	for j := int(0); j <= W; j++ { // 初始化 dp[n][j]
		dp[row][j] = 0
	}
}

// 动态规划
func solve() {
	initDp(n)

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

// dp 正向推导
func solve2() {
	initDp(0)

	for i := 1; i <= n; i++ {
		for j := 0; j <= W; j++ {
			if j < w2[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i-1][j-w2[i]]+v2[i])))
			}
		}
	}

	fmt.Println(dp)
}
func main() {
	//initdp()
	//res := rec(0, W)
	//fmt.Println(res)

	solve2()
}
