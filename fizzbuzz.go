package fizzbuzz

import "strconv"

// FizzBuzz takes an input integer and if
// that divides by 3 returns 'fizz', by 5
// returns 'buzz', by both 3 and 5 returns
// 'fizzbuzz', otherwise the input number.
func FizzBuzz(input int) string {
	if input%5 == 0 && input%3 == 0 {
		return "fizzbuzz"
	}
	if input%5 == 0 {
		return "buzz"
	}
	if input%3 == 0 {
		return "fizz"
	}

	return strconv.Itoa(input)
}
