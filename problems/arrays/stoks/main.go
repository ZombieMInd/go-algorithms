package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 100, 1},
			want:   96,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{1, 2},
			want:   1,
		},
	}

	for _, c := range cases {
		got := maxProfit(c.prices)

		fmt.Println(got)
	}
}

const minMax = -100000

func maxProfit(prices []int) int {
	min, max, dif := prices[0], minMax, 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			max = minMax
		} else if prices[i] > max {
			max = prices[i]
			if max-min > dif {
				dif = max - min
			}
		}
	}

	if dif > 0 {
		return dif
	}
	return 0
}
