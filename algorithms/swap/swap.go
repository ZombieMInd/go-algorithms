package swap

import "fmt"

func Swap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func Test() {
	a := 1
	b := 99

	a, b = Swap(a, b)

	fmt.Println(a)
	fmt.Println(b)
}
