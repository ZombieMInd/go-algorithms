package main

import "fmt"

func main() {
	cases := []struct {
		tree   []*int
		expect int
	}{
		{
			tree: []*int{},
		},
	}

	for _, c := range cases {
		fmt.Printf("got: %v \n expect: %v \n", maxDepth(&TreeNode{}), c.expect)
	}
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return 0
}
