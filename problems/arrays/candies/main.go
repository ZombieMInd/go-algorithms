package main

import "fmt"

func main() {
	cases := []struct {
		ratings []int
		expect  int
	}{
		{
			ratings: []int{1, 0, 2},
			expect:  5,
		},
		{
			ratings: []int{1, 2, 2},
			expect:  4,
		},
		{ // 			  {1, 2, 1, 2, 3, 3, 2, 1}
			ratings: []int{1, 2, 2, 3, 4, 4, 3, 1},
			expect:  15,
		},
		{ //[]int          {1, 2, 3,  1,  3,  2, 1},
			ratings: []int{1, 2, 87, 87, 87, 2, 1},
			expect:  13,
		},
		{ // 			   1,2,3,2,1
			ratings: []int{1, 2, 3, 1, 0},
			expect:  9,
		},
		{ // 			   1, 2, 5,  4, 3, 2, 1
			ratings: []int{1, 6, 10, 8, 7, 3, 2},
			expect:  18,
		},

		{ //   		       1, 2, 3
			ratings: []int{1, 2, 3},
			expect:  6,
		},
	}

	for _, c := range cases {
		fmt.Printf("got: %v \n expect: %v \n", candy(c.ratings), c.expect)
	}
}

func candy(ratings []int) int {
	res := make([]int, len(ratings))

	if len(ratings) == 0 {
		return 0
	}

	if len(ratings) == 1 {
		return 1
	}

	if ratings[0] > ratings[1] {
		res[0] = 2
	} else {
		res[0] = 1
	}

	if ratings[len(ratings)-1] > ratings[len(ratings)-2] {
		res[len(ratings)-1] = 2
	} else {
		res[len(ratings)-1] = 1
	}

	for i := 1; i < len(ratings)-1; i++ {
		res[i] = 1

		if ratings[i] > ratings[i-1] {
			res[i] = res[i-1] + 1
			continue
		}

		if ratings[i] > ratings[i+1] {
			res[i]++
			for j := i; ratings[j] < ratings[j-1] && res[j] >= res[j-1] && j > 1; j-- {
				res[j-1]++
			}
			continue
		}
	}

	if ratings[0] > ratings[1] {
		res[0] = res[1] + 1
	}

	if ratings[len(ratings)-1] > ratings[len(ratings)-2] {
		res[len(ratings)-1] = res[len(ratings)-2] + 1
	}

	c := 0

	for i := range res {
		c += res[i]
	}

	return c
}
