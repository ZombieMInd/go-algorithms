package main

import "fmt"

func main() {
	cases := []struct {
		intervals   [][]int
		newInterval []int
		expect      [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expect:      [][]int{{1, 5}, {6, 9}},
		},

		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expect:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{2, 7},
			expect:      [][]int{{1, 7}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			expect:      [][]int{{1, 5}, {6, 8}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{0, 0},
			expect:      [][]int{{0, 0}, {1, 5}},
		},
		{
			intervals:   [][]int{{0, 0}},
			newInterval: []int{1, 5},
			expect:      [][]int{{0, 0}, {1, 5}},
		},
	}

	for _, c := range cases {

		fmt.Printf("got:\t%v \nexpect:\t%v \n", insert(c.intervals, c.newInterval), c.expect)
	}
}

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	beginInterval := -1

	for i := 0; i < len(intervals); i++ {
		if newInterval[0] >= intervals[i][0] && newInterval[0] <= intervals[i][1] {
			beginInterval = i
		}

		if newInterval[0] < intervals[i][0] && beginInterval == -1 {
			beginInterval = i
		}

		if newInterval[1] >= intervals[i][0] && newInterval[1] <= intervals[i][1] {
			return formOverlap(intervals, newInterval, beginInterval, i, intervals[i][1], 0)
		}

		if newInterval[1] < intervals[i][0] && beginInterval != -1 {
			return formOverlap(intervals, newInterval, beginInterval, i, newInterval[1], 1)
		}

		if newInterval[1] > intervals[i][1] && i == len(intervals)-1 && beginInterval != -1 {
			return formOverlap(intervals, newInterval, beginInterval, i, newInterval[1], 0)
		}
	}

	intervals = append(intervals, newInterval)

	return intervals
}

func formOverlap(intervals [][]int, newInterval []int, beginIndex, currentIndex, endValue, decrement int) [][]int {
	overlapB := intervals[beginIndex][0]
	if newInterval[0] < intervals[beginIndex][0] {
		overlapB = newInterval[0]
	}

	overlap := []int{overlapB, endValue}
	currentIndex -= decrement
	begin := make([][]int, len(intervals[:beginIndex]), len(intervals[:beginIndex])+1) // intervals[:beginIndex]
	copy(begin, intervals[:beginIndex])
	begin = append(begin, overlap)
	end := make([][]int, len(intervals[currentIndex+1:])) // intervals[currentIndex+1:]
	copy(end, intervals[currentIndex+1:])

	return append(begin, end...)
}
