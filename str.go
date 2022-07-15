package exercise

import "fmt"

//strstr("")

// for string
//    for sub
//         i == j; i++ j++
//         i != j;   break;
//    i = 0  j= 0
func main() {
	var str, sub string
	fmt.Scanln(&str, &sub)

	res := strstr(str, sub)
	fmt.Println(res)
}

func strstr(str, sub string) int {
	if str == "" || sub == "" {
		return -1
	}

	for i := 0; i < len(str); i++ {
		pi := i
		for j := 0; j < len(sub); j++ {
			if pi < len(str) && str[pi] == sub[j] {
				pi++
				j++
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
