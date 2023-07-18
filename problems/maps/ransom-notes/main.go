package main

import "fmt"

func main() {
	cases := []struct {
		ransomNote string
		magazine   string
		expect     bool
	}{
		{
			ransomNote: "aa",
			magazine:   "aaab",
			expect:     true,
		},
		{
			ransomNote: "a",
			magazine:   "b",
			expect:     false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			expect:     false,
		},
		{
			ransomNote: "aaa",
			magazine:   "aa",
			expect:     false,
		},
	}

	for _, c := range cases {
		fmt.Printf("got: %v \n expect: %v \n", canConstruct(c.ransomNote, c.magazine), c.expect)
	}
}

func canConstruct(ransomNote string, magazine string) bool {
	letters := map[byte]int{}
	for i := range magazine {
		v, ok := letters[magazine[i]]
		if !ok {
			letters[magazine[i]] = 1
			continue
		}

		letters[magazine[i]] = v + 1
	}

	for i := range ransomNote {
		v, ok := letters[ransomNote[i]]
		if !ok {
			return false
		}

		if v == 0 {
			return false
		}

		letters[ransomNote[i]] = v - 1
	}

	return true
}
