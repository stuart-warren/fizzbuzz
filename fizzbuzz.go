package fizzbuzz

import "strconv"

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
