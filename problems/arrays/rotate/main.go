package main

import "fmt"

func main() {
	cases := []struct {
		nums   []int
		steps  int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7},
			steps:  8,
			expect: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums:   []int{-1, -100, 3, 99},
			steps:  2,
			expect: []int{3, 99, -1, -100},
		},
		{
			nums:   []int{1, 2},
			steps:  5,
			expect: []int{2, 1},
		},
	}

	for _, c := range cases {
		rotateMem(c.nums, c.steps)
		fmt.Printf("got: %v \n expect: %v \n", c.nums, c.expect)
	}
}

// func rotate(nums []int, k int) {
// 	buf := 0
// 	current := 0

// 	for i := 0; i < k; i++ {
// 		buf = nums[len(nums)-1]
// 		prev := nums[0]

// 		for j := 1; j < len(nums); j++ {
// 			current = nums[j]
// 			nums[j] = prev
// 			prev = current
// 		}

// 		nums[0] = buf
// 	}
// }

func rotateMem(nums []int, k int) {
	if k > len(nums) {
		k = k % len(nums)
	}

	begin := make([]int, k)
	end := make([]int, len(nums)-k)

	copy(begin, nums[len(nums)-k:])
	copy(end, nums[:len(nums)-k])

	res := append(begin, end...)

	copy(nums, res)
}
