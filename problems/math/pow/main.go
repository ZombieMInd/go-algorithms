package main

import (
	"fmt"
	"math"
)

// Unsolved

func main() {
	cases := []struct {
		x      float64
		n      int
		expect float64
	}{
		// {
		// 	x:      2,
		// 	n:      10,
		// 	expect: 1024,
		// },
		// {
		// 	x:      2.1,
		// 	n:      3,
		// 	expect: 9.261,
		// },
		// {
		// 	x:      2,
		// 	n:      -2,
		// 	expect: 0.25,
		// },
		// {
		// 	x:      1,
		// 	n:      -2147483648,
		// 	expect: 1,
		// },
		// {
		// 	x:      2,
		// 	n:      -2147483648,
		// 	expect: 1,
		// },
		// {
		// 	x:      -1,
		// 	n:      -2147483648,
		// 	expect: 1,
		// },
		// {
		// 	x:      -1,
		// 	n:      -2147483647,
		// 	expect: 1,
		// },
		{
			x:      1.0000000000001,
			n:      -2147483648,
			expect: 1,
		},
	}

	for _, c := range cases {
		res := myPow(c.x, c.n)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, math.Pow(c.x, float64(c.n)))
		// if res != c.expect {
		// 	panic("not equal")
		// }
	}
}

func myPow(x float64, n int) float64 {
	res := 1.0

	if x == 1 {
		return 1
	}

	if x == 0 && n != 0 {
		return 0
	}

	if x == -1 {
		if n%2 == 0 {
			return 1
		} else {
			return -1
		}
	}

	if n > 0 {
		for i := 0; i < n; i++ {
			res *= x
		}
	} else if n < 0 {
		for i := 0; i < -n && i < 10000000; i++ {
			res /= x
			if res == 0 {
				return res
			}
		}
	}

	return res
}
