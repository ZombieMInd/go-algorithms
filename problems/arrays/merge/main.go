package main

import "fmt"

func main() {
	cases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
		},
		{
			nums1: []int{1, 0},
			m:     1,
			nums2: []int{2},
			n:     1,
		},
		{
			nums1: []int{2, 0},
			m:     1,
			nums2: []int{1},
			n:     1,
		},
	}

	for _, c := range cases {
		merge(c.nums1, c.m, c.nums2, c.n)
		fmt.Println(c.nums1)
	}
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	k := 0
	for i := 0; i < n+m; i++ {
		if nums2[k] <= nums1[i] {
			for j := m + k - 1; j >= i; j-- {
				nums1[j+1] = nums1[j]
			}

			nums1[i] = nums2[k]
			k++
			if k == n {
				break
			}
		}
	}

	if k+1 <= n {

		for i := 0; i < n-k; i++ {
			nums1[m+k+i] = nums2[i+k]
		}
	}
}

// for k := m; k >= i; k-- {
//     nums1[k+1] = nums1[k]
// }
