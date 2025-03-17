package fizzbuzz

import (
	"strconv"
)

func FizzBuzz(i int) (res string) {
	if i%15 == 0 {
		res = "FizzBuzz"
	} else if i%5 == 0 {
		res = "Buzz"
	} else if i%3 == 0 {
		res = "Fizz"
	} else {
		res = strconv.Itoa(i)
	}
	return
}
