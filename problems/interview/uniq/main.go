// Назовем натуральное число "особым", если все его цифры различны.
// Для заданного положительного целого числа n верните количество особых целых чисел,
// принадлежащих интервалу [1, n].
//
// Пример
// Аргумент: n = 20
// Вывод: 19
// Обьяснение: Все числа от 1 до 20, кроме 11, являются особыми.
// Таким образом, существует 19 особых целых чисел.
//
// Ограничения
// 1 <= n <= 2 * 10^9

// 200

package main

import (
	"fmt"
)

func countSpecialNumbersOld(n int) int {
	result := 0

numbersLoop:
	for i := 1; i <= n; i++ {
		num := i
		wordDigits := map[int]struct{}{}

		for num > 0 {
			digit := num % 10
			num /= 10

			_, ok := wordDigits[digit]
			if ok {
				continue numbersLoop
			}

			wordDigits[digit] = struct{}{}
		}

		result++
	}

	return result
}

func countSpecialNumbersRange(m, n int) int {
	result := 0

numbersLoop:
	for i := m; i <= n; i++ {
		num := i
		wordDigits := map[int]struct{}{}

		for num > 0 {
			digit := num % 10
			num /= 10

			_, ok := wordDigits[digit]
			if ok {
				continue numbersLoop
			}

			wordDigits[digit] = struct{}{}
		}

		result++
	}

	return result
}

// func countSpecialNumbers(n int) int {
// 	res := 0
// 	num := n
// 	digits := []int{}
// 	for num > 0 {
// 		digits = append(digits, num%10)
// 		num /= 10
// 	}

// 	res = (digits[len(digits)-1] + 1) * A(9, len(digits))
// 	res = res - countSpecialNumbersRange(n, (digits[len(digits)-1]+1)*int(math.Pow(10, float64(len(digits)-1)))-1)
// 	return res
// }

func digitCombinations(n int) int {
	return A(10, n)
}

func A(n, m int) int {
	return factorial(n) / (factorial(n - m))
}

func C(n, m int) int {
	return A(n, m) / factorial(m)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// type TestCase struct {
	// 	Given    int
	// 	Expected int
	// }

	// tCases := []TestCase{
	// 	// 		{1, 1},
	// 	// 		{5, 5},
	// 	// {20, 19},     // 1
	// 	// {30, 28},     // 2
	// 	// {50, 46},     // 4
	// 	// {100, 90},    // 10
	// 	// {120, 99},    // 21
	// 	// {150, 123},   // 27
	// 	// {200, 162},   // 38
	// 	// {250, 1},     //
	// 	// {2000, 1242}, // 758
	// 	// 		{123456789, 2392084},
	// {99, 90},
	// }

	// n / 10 + n / 100 * 9
	// + n / 1000 * 9 * 8

	// 90 + 81 = 171
	// 234 - 171 = 63
	// 24 + 18 = 42
	// 23 + 2*9 = 41
	// 306 - 234 = 72
	// 234 - 162 = 72

	// for i := 14; i <= 654; i += 1 {
	// 	pr(i)
	// }

	// pr(654)
	// pr(54)

	// 4 + 5*10 - (5-1) = 50
	// 4 + 5 * 10 - (5-1) + 6 * 100 - (6*)

	// for i := 20000; i <= 70000; i += 1000 {
	// 	fmt.Println(countSpecialNumbersOld(i) - countSpecialNumbersOld(i-1000))
	// }

	// for i := 20000; i <= 70000; i += 10000 {
	// 	fmt.Println(countSpecialNumbersOld(i) - countSpecialNumbersOld(i-10000))
	// }

	// fmt.Println(countSpecialNumbersOld(1000))
	// fmt.Println(countSpecialNumbersOld(2000))
	// fmt.Println(countSpecialNumbers2(20))
	// fmt.Println(countSpecialNumbers2(30))
	// fmt.Println(countSpecialNumbersOld(150), countSpecialNumbers(150))
	fmt.Println(countSpecialNumbers(99))
	fmt.Println(countSpecialNumbers(30))
	// fmt.Println(countSpecialNumbers(150))
	fmt.Println(countSpecialNumbers(123456789))
	// for _, tCase := range tCases {
	// 	given := countSpecialNumbers(tCase.Given)
	// 	fmt.Println(
	// 		"given", tCase,
	// 		"diff", tCase.Given-given,
	// 		"result-formula", (tCase.Given-given)-(tCase.Given/10-tCase.Given%(tCase.Given/10)+1+tCase.Given/10*tCase.Given/100),
	// 		"n/10-n%(n/10)+1+n/10*n/100", tCase.Given/10-tCase.Given%(tCase.Given/10)+tCase.Given/10*tCase.Given/100,
	// 	)
	// 	// if given != tCase.Expected {
	// 	// 	println("Test case " + strconv.Itoa(i) + " failed!")
	// 	// } else {
	// 	// 	println("Test case " + strconv.Itoa(i) + " passed!")
	// 	// }
	// }
}

func countSpecialNumbers(n int) int {
	if n < 10 {
		return n
	}

	num := n
	digits := []int{}
	for num > 0 {
		digits = append([]int{num % 10}, digits...)
		num /= 10
	}

	res := 0

	for i := 1; i < len(digits); i++ {
		x := 1
		k := 9
		for j := 0; j < i-1; j++ {
			x *= k
			k--
		}
		res += 9 * x
	}

	takenDigits := map[int]struct{}{}
	for i := 0; i < len(digits); i++ {
		choices := 0
		for j := 0; j < digits[i]; j++ {
			if _, ok := takenDigits[j]; !ok {
				choices++
			}
		}
		if i == 0 && digits[i] > 0 {
			choices--
		}
		begin := 1
		remains := 10 - i - 1
		for j := i + 1; j < len(digits); j++ {
			begin *= remains
			remains--
		}
		res += choices * begin

		if _, ok := takenDigits[digits[i]]; ok {
			return res
		}
		takenDigits[digits[i]] = struct{}{}

	}

	return res + 1
}
