package main

//1. rand13 - rand5
//1 - 13，对其模6，其中
//（1 2 3 4 5 、 7 8 9 10 11 、 13 ）% 6 = （1 - 5）
//为了保证分布均匀，13舍去
func rand13() int{
	return -1
}

func rand5() int{
	for{
		r := rand13()
		if r % 6 == 0 || r == 13{
			continue
		}
		return r % 6
	}
}


//2. rand5 - rand13
//首先推算需要多次使用rand5，将其值做加或者乘运算，满足 13+ 的值域
//加法：因为 1 + n > 1  (n为正整数) ，不合适，所以需要混合运算
//同时需要保证分布均匀
//[（1 - 5 ) - 1 ]  =  （0 - 4 ） * 5 = （0 5 10 15 20 ） + rand5
//值域为（ 1 2 3 4 5 、 6 7 8 9 10、 11 12 13 14 15、 16 17 18 19 20 ），舍弃（14 - 20）
func _rand5() int{
	return -1
}
func _rand13() int{
	for{
		r := rand5() + (rand5() - 1) * 5
		if r > 13{
			continue
		}
		return r
	}
}