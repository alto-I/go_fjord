package fizzbuzz

import "strconv"

func Fizzbuzz(i int) string {
	if i%15 == 0 {
		return "fizzbuzz"
	} else if i%5 == 0 {
		return "buzz"
	} else if i%3 == 0 {
		return "fizz"
	} else {
		return strconv.Itoa(i)
	}
}
