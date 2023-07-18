package main

import (
	"fmt"
	"strings"
)

func main() {
	cases := []struct {
		s      string
		expect bool
	}{
		{
			s:      "A man, a plan, a canal: Panama",
			expect: true,
		},
		{
			s:      "race a car",
			expect: false,
		},
		{
			s:      " ",
			expect: true,
		},
		{
			s:      ".a.",
			expect: true,
		},
		{
			s:      "!!!::vrTTr::v!!!",
			expect: true,
		},
	}

	for _, c := range cases {
		fmt.Printf("got:\t%v \nexpect:\t%v \n", isPalindrome(c.s), c.expect)
	}
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	defer func() {
		fmt.Println(s)
	}()

	for i := 0; i < len(s); i++ {
		if !isAlfaNumeric(int(s[i])) {
			s = s[:i] + s[i+1:]
			i--
		}
	}

	for i := 0; i < len(s)/2; i++ {

		if int(s[i]) != int(s[len(s)-1-i]) {
			return false
		}

	}

	return true
}

func isAlfaNumeric(code int) bool {
	return (code >= 48 && code <= 57) || (code >= 97 && code <= 122)
}
