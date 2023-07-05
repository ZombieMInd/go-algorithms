package main

// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// Given two strings needle and haystack,
// return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.

func main() {
	testCases := []struct {
		haystack string
		needle   string
		want     int
	}{
		{
			haystack: "hello",
			needle:   "ll",
			want:     2,
		},
		{
			haystack: "aaaaa",
			needle:   "bba",
			want:     -1,
		},
		{
			haystack: "a",
			needle:   "a",
			want:     0,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			want:     4,
		},
	}

	for _, tc := range testCases {
		got := strStr(tc.haystack, tc.needle)

		if got != tc.want {
			panic("wrong")
		}
	}
}

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}
