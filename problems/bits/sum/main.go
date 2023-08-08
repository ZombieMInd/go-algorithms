package main

import (
	"fmt"
)

func main() {
	cases := []struct {
		a, b   string
		expect string
	}{
		// {
		// 	a:      "11",
		// 	b:      "1",
		// 	expect: "100",
		// },
		{
			a:      "1010",
			b:      "1011",
			expect: "10101",
		},
		{
			a:      "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			b:      "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			expect: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for _, c := range cases {
		res := addBinary(c.a, c.b)
		fmt.Printf("got:\t%v \nexpect:\t%v \n", res, c.expect)
		if res != c.expect {
			panic("not equal")
		}
	}
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return sum(a, b)
	} else {
		return sum(b, a)
	}
}

func sum(a, b string) string {

	res := ""
	remain := '0'
	for i := 0; i < len(a); i++ {
		if i >= len(b) {
			if remain == '1' {
				if a[i] == '1' {
					res += "0"
				} else {
					res += "1"
					remain = '0'
				}
			}
		} else {
			if a[i] == '1' && b[i] == '1' {
				if remain == '1' {
					res += "1"
				} else {
					res += "0"
					remain = '1'
				}
			} else if a[i] == '1' || b[i] == '1' {
				if remain == '1' {
					res += "0"
				} else {
					res += "1"
					remain = '0'
				}
			} else {
				if remain == 1 {
					res += "1"
					remain = '0'
				}
			}
		}
	}

	if res[len(res)-1] == '0' {
		if remain != '0' {
			res += "1"
		} else {
			res = res[:len(res)-1]
		}
	}

	resRevers := make([]byte, len(res))

	for i := range res {
		resRevers[len(res)-1-i] = res[i]
	}

	return string(resRevers)
}
