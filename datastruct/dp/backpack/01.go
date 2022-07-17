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





// INF
// 大重量 O(nW)复杂度太高
// 1 <= n <= 100        MAXN = 100
// 1 <= wi <= 10^7
// 1 <= vi <= 100		MAXV = 100
// 1 <= W <= 10^9
const (
	INF  = 0x8fffffff
	MAXS = 50
)

// dp[i][j]    0 <= j <= MAXN * MAXV + 1
var (
	sdp [MAX][MAXS + 1]int
)

func initsdp() {
	sdp[0][0] = 0
	for j := 1; j <= MAXS; j++ {
		sdp[0][j] = 99 //math.MaxInt
	}
}
func solve3() {
	initsdp()

	for i := 0; i < n; i++ {
		for j := 0; j <= MAXS; j++ {
			if j < v[i] {
				sdp[i+1][j] = sdp[i][j]
			} else {
				sdp[i+1][j] = int(math.Min(float64(sdp[i][j]), float64(sdp[i][j-v[i]]+w[i])))
			}
		}
	}

	res := 0
	for j := 0; j <= MAXS; j++ {
		if sdp[n][j] <= W {
			res = j
		}
	}
	fmt.Println(res)
	//for i := 0; i <= n; i++ {
	//	fmt.Println(i, sdp[i])
	//}
}


func main() {
	//initdp()
	//res := rec(0, W)
	//fmt.Println(res)

	//solve2()
	solve3()
}
