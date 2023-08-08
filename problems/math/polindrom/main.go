package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		x      int
		expect bool
	}{
		{
			x:      121,
			expect: true,
		},
		{
			x:      -121,
			expect: false,
		},
		{
			x:      10,
			expect: false,
		},
	}

	for _, c := range cases {
		res := isPalindrome(c.x)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, c.expect)
		// if res != c.expect {
		// 	panic("not equal")
		// }
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	strX := fmt.Sprintf("%d", x)

	for i := 0; i < len(strX)/2; i++ {
		if strX[i] != strX[len(strX)-1-i] {
			return false
		}
	}

	return true
}
