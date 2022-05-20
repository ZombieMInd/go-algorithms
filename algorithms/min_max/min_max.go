package min_max

import "fmt"

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Test() {
	fmt.Println(Max(1, 100))
	fmt.Println(Min(1, 100))
}
