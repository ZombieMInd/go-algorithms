package main

import (
	"fmt"
)

//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]

// ["1","1","0","0","0"],
// ["1","1","0","0","0"],
// ["0","0","1","0","0"],
// ["0","0","0","1","1"]

func main() {
	cases := []struct {
		grid   [][]byte
		expect int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expect: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expect: 3,
		},
	}

	for _, c := range cases {

		fmt.Printf("got:\t%v \nexpect:\t%v \n", numIslands(c.grid), c.expect)
	}
}

type Point struct {
	i, j int
}

func NewPoint(i, j int) Point {
	return Point{i: i, j: j}
}

func numIslands(grid [][]byte) int {

	islandsCounter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				BFS(grid, i, j)
				grid[i][j] = 0
				islandsCounter++
			}
		}
	}

	return islandsCounter
}

func BFS(grid [][]byte, beginI, beginJ int) {
	q := NewQueue()

	getNearestNeighbors(grid, beginI, beginJ, q)

	for node := q.Dequeue(); node != nil; node = q.Dequeue() {
		getNearestNeighbors(grid, node.i, node.j, q)
	}
}

func getNearestNeighbors(grid [][]byte, i, j int, q *Queue) {
	if i-1 >= 0 {
		if grid[i-1][j] == '1' {
			q.Enqueue(NewPoint(i-1, j))
			grid[i-1][j] = 0
		}
	}

	if j-1 >= 0 {
		if grid[i][j-1] == '1' {
			q.Enqueue(NewPoint(i, j-1))
			grid[i][j-1] = 0
		}
	}

	if i+1 < len(grid) {
		if grid[i+1][j] == '1' {
			q.Enqueue(NewPoint(i+1, j))
			grid[i+1][j] = 0
		}
	}

	if j+1 < len(grid[0]) {
		if grid[i][j+1] == '1' {
			q.Enqueue(NewPoint(i, j+1))
			grid[i][j+1] = 0
		}
	}
}

type Queue []Point

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(v Point) {
	*q = append(*q, v)
}

func (q *Queue) Dequeue() *Point {
	if len(*q) == 0 {
		return nil
	}

	v := (*q)[0]
	*q = (*q)[1:]
	return &v
}
