package main

import "fmt"

const MAX2 = 5 + 5
var (
	s1, s2 string
	dp [MAX2][MAX2]int
)

func solve(){
	m := len(s1)
	n := len(s2)
	maxLen := 0
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			if s1[i - 1] == s2[j - 1]{
				dp[i][j] = dp[i - 1][j - 1] + 1
				if maxLen < dp[i][j]{
					maxLen = dp[i][j]
				}
			}else {
				dp[i][j] = dp[i - 1][j - 1]
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(maxLen)
}
func main() {
	fmt.Scanf("%s\n%s\n", &s1, &s2)
	//fmt.Println("s1:", s1, "s2:", s2)
	solve()
}
