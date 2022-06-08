// https://leetcode.com/problems/reverse-integer/

/*
Explanation:

Boring task with small additional range check, almost same as https://leetcode.com/problems/reverse-integer/

!!! No intresting notes !!!
*/

package main

import "fmt"

func myAtoi(s string) int {
	isNegative := 0
	charCount := 0
	resVal := 0
	sLen := len(s)

	// Skip leading spaces
	for charCount < sLen && s[charCount] == ' ' {
		charCount++
	}

	if charCount == sLen {
		return resVal
	}

	// Check sign
	if s[charCount] == '-' {
		isNegative = 1
		charCount++
	} else if s[charCount] == '+' {
		charCount++
	}

	// While in characters are numbers
	for charCount < sLen && (s[charCount] >= '0' && s[charCount] <= '9') {
		// Conver char to the numeric value
		tmpVal := int(s[charCount] - '0')

		// Check overflow
		if resVal > 214748364-tmpVal {
			if isNegative > 0 {
				return -2147483648
			} else {
				return 2147483647
			}
		}

		resVal = resVal*10 + tmpVal
		charCount++
	}

	if isNegative > 0 {
		resVal *= -1
	}

	return resVal
}

func main() {

	fmt.Printf("42: %d\n", myAtoi("42"))
	fmt.Printf("   -42: %d\n", myAtoi("   -42"))
	fmt.Printf("4193 with words: %d\n", myAtoi("4193 with words"))

}
