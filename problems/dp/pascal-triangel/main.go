package main

import "fmt"

func main() {
	cases := []struct {
		numRows int
		expect  [][]int
	}{
		{
			numRows: 6,
			expect:  [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}},
		},
	}

	for _, c := range cases {
		fmt.Printf("got: %v \n expect: %v \n", generate(c.numRows), c.expect)
	}
}

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		generateRow(res, i)
	}

	return res
}

func generateRow(res [][]int, row int) {
	res[row] = make([]int, row+1)

	res[row][0] = 1

	for i := 1; i < row; i++ {
		res[row][i] = res[row-1][i-1] + res[row-1][i]
	}

	res[row][len(res[row])-1] = 1
}
