package main

import "fmt"

const MAX = 50 + 5
var (
	f [MAX][MAX][MAX][MAX]int
	dp [MAX][MAX]int
	m, n int
)

func solve() int{
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			for k := m; k >= 1; k--{
				for l := n; l >= 1; l--{
					f[i][j][k][l] = _max(f[i - 1][j][k - 1][l], f[i - 1][j][k][l - 1], f[i][j - 1][k - 1][l], f[i][j - 1][k][l - 1]) + dp[i][j] + dp[k][l]
					if i == k && j == l{
						f[i][j][k][l] -= dp[k][l]
					}
				}
			}
		}
	}
	//fmt.Println(f[m][n][m][n])
	return f[m][n][m][n]
}


func _max(a, b, c, d int) (max int){
	var max1, max2 int
	if a > b{
		max1 = a
	}else{
		max1 = b
	}

	if c > d{
		max2 = c
	}else{
		max2 = d
	}

	if max1 > max2{
		max = max1
	}else{
		max = max2
	}
	return
}
func main() {
	fmt.Scanln(&m, &n)
	var data int
	for i := 1; i <= m; i++{
		for j := 1; j <= n; j++{
			fmt.Scan(&data)
			dp[i][j] = data
		}
		fmt.Scanln()
	}

	res := solve()
	fmt.Print(res)
}
