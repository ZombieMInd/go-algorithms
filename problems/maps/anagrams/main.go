package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		strs   []string
		expect [][]string
	}{
		{
			strs:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expect: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}

	for _, c := range cases {
		res := groupAnagrams(c.strs)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, c.expect)
		// if res != c.expect {
		// 	panic("not equal")
		// }
	}
}

func groupAnagrams(strs []string) [][]string {
	res := [][]int{}
	resMaps := []map[byte]int{}

	for i := range strs {
		for j := range resMaps {
			for l := range strs[i] {

			}
		}
	}
}
