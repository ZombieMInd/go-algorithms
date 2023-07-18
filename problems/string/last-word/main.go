package main

import "fmt"

func main() {
	cases := []struct {
		s      string
		expect int
	}{
		{
			s:      "Hello World",
			expect: 5,
		},
		{
			s:      "   fly me   to   the moon  ",
			expect: 4,
		},
		{
			s:      "luffy is still joyboy",
			expect: 6,
		},
		{
			s:      "luffy is still    j   ",
			expect: 1,
		},
	}

	for _, c := range cases {
		fmt.Println(lengthOfLastWord(c.s))
	}
}

func lengthOfLastWord(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if l > 0 && s[i] == ' ' {
			return l
		}

		if s[i] != ' ' {
			l++
		}
	}

	return l
}
