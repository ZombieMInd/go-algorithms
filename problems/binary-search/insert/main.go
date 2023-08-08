package main

import (
	"fmt"
	"math"
)

func main() {
	cases := []struct {
		nums   []int
		target int
		expect int
	}{
		{
			nums:   []int{1, 3, 5, 6},
			target: 5,
			expect: 2,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 2,
			expect: 1,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 7,
			expect: 4,
		},
		{
			nums:   []int{1, 3, 5, 6},
			target: 0,
			expect: 0,
		},
		{
			nums:   []int{1, 3, 5, 20},
			target: 6,
			expect: 3,
		},
		{
			nums:   []int{1, 3},
			target: 1,
			expect: 0,
		},
		{
			nums:   []int{1},
			target: 3,
			expect: 1,
		},
		{
			nums:   []int{1},
			target: 0,
			expect: 0,
		},
		{
			nums:   []int{1, 3, 5},
			target: 0,
			expect: 0,
		},
		{
			nums:   []int{2, 3, 5, 6, 9},
			target: 7,
			expect: 4,
		},
		{
			nums:   longNums,
			target: 5488,
			expect: 6363,
		},
	}

	for _, c := range cases {
		res := searchInsert(c.nums, c.target)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, c.expect)
		if res != c.expect {
			panic("not equal")
		}
	}
}

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	pos := len(nums) / 2
	div := pos
	for i := 0; float64(i) < math.Log2(float64(len(nums))); i++ {
		if nums[pos] == target {
			return pos
		}

		if nums[pos] > target {
			div = div / 2
			if div == 0 {
				div = 1
			}
			pos -= div
			continue
		}

		if nums[pos] < target {
			div = div / 2
			if div == 0 {
				div = 1
			}
			pos += div
			continue
		}
	}

	if pos < 0 || pos < len(nums) && nums[pos] < target {
		pos++
	}

	if pos > len(nums) || pos > 0 && nums[pos-1] > target {
		pos--
	}

	return pos
}
