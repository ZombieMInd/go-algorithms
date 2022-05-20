package squares

import "fmt"

func Sum(c int) int {
	res := 0
	for i := 1; i <= c; i++ {
		res += i * i
	}

	return res
}

func Test() {
	fmt.Println(Sum(5))
}
