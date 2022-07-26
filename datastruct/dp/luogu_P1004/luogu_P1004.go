package main

// 本来考虑用贪心，分两步走两遍，但是不行，每次的极值加起来不一定是最值
//0 3 5
//2 6 0
//0 4 0
//第一次求最大值的话，是0->3->6->4->0
//第二次求最大值的话，是0->3（已经被去掉了，为了方便看）->5->0->0
//合起来是18，而实际答案应该是20：
//第一次：0->3->5->0->0
//第二次：0->2->6->4->0

import "fmt"

const MAX = 10

var(
	n int
	f [MAX][MAX][MAX][MAX]int
	dp2 [MAX][MAX]int
)

// 计算完成后回溯清空路径
func ff() int{
	// recu
	for i := 1; i <= n; i++{
		for j := 1; j <= n; j++{
			for k := 1; k <= n; k++{
				for l := 1; l <= n; l++{
					f[i][j][k][l] = _max(f[i - 1][j][k - 1][l], f[i - 1][j][k][l - 1], f[i][j - 1][k - 1][l], f[i][j - 1][k][l - 1]) + dp2[i][j] + dp2[k][l]
					if i == k && j == l{
						f[i][j][k][l] -= dp2[k][l]
					}
				}
			}
		}
	}
	
	return f[n][n][n][n]
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
	var i, j, data int
	fmt.Scanln(&n)
	for {
		fmt.Scanln(&i, &j, &data)
		if i == 0{
			break
		}
		dp2[i][j] = data
	}

	res := ff()
	fmt.Println(res)

}