package main

import (
	"fmt"
)

// 1 1 2 3 5 8 13 21 34 55 89 144

func main() {
	cases := []struct {
		n      int
		expect int
	}{
		{
			n:      2,
			expect: 2,
		},
		{
			n:      3,
			expect: 3,
		},
		{
			n:      4,
			expect: 5,
		},
	}

	for _, c := range cases {
		res := climbStairs(c.n)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, c.expect)
		// if res != c.expect {
		// 	panic("not equal")
		// }
	}
}

func climbStairs(n int) int {
	res1 := 1
	res2 := 1
	for i := 1; i < n; i++ {
		res1, res2 = res2, res1+res2
	}

	return res2
}
