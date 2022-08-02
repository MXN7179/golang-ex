package main

import (
	"fmt"
	"math"
)

func isPrime(data int) bool{
	if data <= 1{
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(data))); i++{
		if data % i == 0{
			return false
		}
	}
	return true
}
func factor(data int) {
	for i := 1; i <= int(math.Sqrt(float64(data))); i++{
		if data % i == 0{
			j := data / i
			if isPrime(i) == true{
				fmt.Printf("%d ", i)
			}
			if isPrime(j) == true{
				fmt.Printf("%d ", j)
			}
		}
	}
	fmt.Println()
}
func main() {

	for{
		c := 0
		fmt.Scanf("%d\n", &c)
		//fmt.Println(isPrime(c))
		factor(c)
	}
}