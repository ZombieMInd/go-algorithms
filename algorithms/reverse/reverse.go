package reverse

import "fmt"

func Reverse(s []int) []int {
	for a, b := 0, len(s)-1; a < b; a, b = a+1, b-1 {
		s[a], s[b] = s[b], s[a]
	}
	return s
}

func Test() {
	res := Reverse([]int{1, 2, 3, 4, 5})
	fmt.Println(res)
}
