package main

import (
	"fmt"
)

const(
	MAX = 20 + 1
)
var(
	p, q int
	n, m int
	dp [MAX][MAX]int
)

func Init(){
	// init
	dp[0][0] = 0
	for i := 1; i <= n; i++{
		if valid(i, 0) == false{
			//fmt.Printf("not vaild i:%d j:0\n", i )
			break
		}
		dp[i][0] = 1
	}
	for j := 1; j <= m; j++{
		if valid(0, j) == false{
			//fmt.Printf("not vaild i:0 j:%d\n", j )
			break
		}
		dp[0][j] = 1
	}
}

func solve() int{

	for i := 1; i <= n; i++{
		for j := 1; j <= m; j++{
			cur := 0
			if valid(i, j - 1) == true{
				cur += dp[i][j - 1]
			}
			if valid(i - 1, j) == true{
				cur += dp[i - 1][j]
			}
			dp[i][j] = cur
		}
	}
	return dp[n][m]
}

func print(){
	for i := 0; i <= n; i++{
		for j := 0; j <= m; j++{
			fmt.Printf("%3d ", dp[i][j])
		}
		fmt.Println()
	}
}

func valid(i, j int) bool{
	//fmt.Printf("valid i: %v j: %v\n", i, j)
	if !(i >= 0 && j >= 0){
		return false
	}else if i == p && j == q {
		return false
	}else if i == p + 1 && (j == q - 2 || j == q + 2){
		return false
	}else if i == p - 1 && (j == q - 2 || j == q + 2){
		return false
	}else if i == p + 2 && (j == q - 1 || j == q + 1){
		return false
	}else if i == p - 2 && (j == q - 1 || j == q + 1){
		return false
	}
	return true
}

func main() {
	fmt.Scanln(&n, &m, &p, &q)
	Init()
	res := solve()
	fmt.Println(res)
	//print()
}