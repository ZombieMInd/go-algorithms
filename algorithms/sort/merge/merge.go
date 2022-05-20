package merge

import "fmt"

func Sort(s []int) []int {
	if len(s) == 1 {
		return s
	}

	n := len(s)/2 + len(s)%2

	done := make(chan bool)

	var left []int
	go func() {
		left = Sort(s[:n])
		done <- true
	}()

	right := Sort(s[n:])

	<-done

	return sort(left, right)
}

func sort(left, right []int) []int {
	var res []int
	var l, r int

	for l, r = 0, 0; l < len(left) && r < len(right); {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	res = append(res, left[l:]...)
	res = append(res, right[r:]...)

	return res
}

func Test() {
	fmt.Println(Sort([]int{10, -1, 2, 5, 0, 6, 4, -999, -5}))
}
