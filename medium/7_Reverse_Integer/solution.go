// https://leetcode.com/problems/reverse-integer/

/*
Explanation:

Boring task with small additional range check

!!! No intresting notes !!!
*/

package main

import "fmt"

func reverse(x int) int {
	resVal := 0

	isNegative := x < 0
	if isNegative {
		x *= -1
	}

	for x > 0 {
		if resVal+(x%10) > 214748365 {
			return 0
		}

		resVal *= 10
		resVal += x % 10
		x /= 10
	}

	if isNegative {
		resVal *= -1
	}
	return resVal
}

func main() {

	fmt.Printf("123: %d\n", reverse(123))
	fmt.Printf("-123: %d\n", reverse(-123))
	fmt.Printf("120: %d\n", reverse(120))

}
