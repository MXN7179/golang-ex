package main

import "fmt"

//strstr("")

// for string
//    for sub
//         i == j; i++ j++
//         i != j;   break;
//    i = 0  j= 0
func main() {
	for {
		var str, sub string
		fmt.Println("in:")
		fmt.Scanln(&str)
		fmt.Scanln(&sub)

		res := strstr(str, sub)
		fmt.Println(res)
	}

}

func strstr(str, sub string) int {
	if str == "" || sub == "" {
		return -1
	}

	for i := 0; i < len(str); i++ {
		pi := i
		j := 0
		for ; j < len(sub); j++ {
			if pi < len(str) && str[pi] == sub[j] {
				pi++
				//j++
			} else {
				break
			}
		}
		if j >= len(sub) {
			return i
		}
	}

	return -1
}
