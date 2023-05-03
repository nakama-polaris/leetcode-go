package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(romanToInt(s))
}

func romanToInt(s string) int {
	ans := 0

	var romanInt = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	var specialRomanInt = map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}

	for cursor := 0; cursor < len(s); {
		if cursor == (len(s) - 1) {
			ans += romanInt[s[cursor:cursor+1]]
			cursor += 1
		} else if romanInt[s[cursor:cursor+1]] >= romanInt[s[cursor+1:cursor+2]] {
			ans += romanInt[s[cursor:cursor+1]]
			cursor += 1
		} else {
			ans += specialRomanInt[s[cursor:cursor+2]]
			cursor += 2
		}
	}
	return ans
}
