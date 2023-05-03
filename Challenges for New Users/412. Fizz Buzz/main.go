package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 20
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	answer := make([]string, n)

	// while the number reaches n
	for i := 1; i <= n; i++ {
		// if the number is divisible by 3 and 5, put "FizzBuzz" in the answer
		if i%3 == 0 && i%5 == 0 {
			answer[i-1] = "FizzBuzz"
			// if the number is divisible by 3, put "Fizz" in the answer
		} else if i%3 == 0 {
			answer[i-1] = "Fizz"
			// if the number is divisible by 5, put "Buzz" in the answer
		} else if i%5 == 0 {
			answer[i-1] = "Buzz"
			// when none of the above is true, put the number as string in the answer
		} else {
			answer[i-1] = strconv.Itoa(i)
		}
	}

	return answer
}
