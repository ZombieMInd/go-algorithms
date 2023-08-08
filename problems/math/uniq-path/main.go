package main

import (
	"fmt"
	"math/big"
	"time"
)

// 2 - 1 turn
// 5 + 1 - 2 turns
// 5 + 5 - 3 turns
// 4 + 3 + 2 + 1 - 4 turns

func main() {
	fmt.Println(time.Now().AddDate(0, -1, -1))
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}

	if m > n {
		m, n = n, m
	}

	return int(big.NewInt(0).Div(factorial(m+n-2), big.NewInt(0).Mul(factorial(m-1), factorial(n-1))).Int64())
}

func factorial(n int) *big.Int {
	res := big.NewInt(1)
	for i := 1; i <= n; i++ {
		res = res.Mul(res, big.NewInt(int64(i)))
	}

	return res
}
