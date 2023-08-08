package main

import (
	"fmt"
)

// Unsolved

func main() {
	cases := []struct {
		x      []int
		expect bool
	}{
		{
			x:      []int{1, 2, 3, 1},
			expect: true,
		},
		{
			x:      []int{1, 2, 3, 4},
			expect: false,
		},
	}

	for _, c := range cases {
		res := containsDuplicate(c.x)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, c.expect)
		// if res != c.expect {
		// 	panic("not equal")
		// }
	}
}

func containsDuplicate(nums []int) bool {
	uniqNums := map[int]struct{}{}

	for i := 0; i < len(nums); i++ {
		_, ok := uniqNums[nums[i]]
		if ok {
			return true
		}

		uniqNums[nums[i]] = struct{}{}
	}

	return false
}
