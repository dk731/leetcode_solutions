// https://leetcode.com/problems/palindrome-number/

/*
Explanation:

A better approach is available by creating inverted number,
and returning if it is equal to the original (with this approach
even bigger shortcut can be made, by reversing only half of the
digits). This solution is taking approach very simillar to the
5. longest palindromic substring, by comparing extracting numbers
one by one and comparing them

!!! Use precalculated tenth table !!!
*/

package main

import "fmt"

var tenthPow [16]int = [16]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000, 10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000}

func isPalindrome(x int) bool {
	// Exclude nagative numbers from the very beggining
	if x < 0 {
		return false
	}

	digitsCount := 1

	for x/tenthPow[digitsCount] >= 1 {
		digitsCount++
	}

	rp := digitsCount - 1
	lp := 0

	for lp < rp {
		if (x/tenthPow[lp])%10 != (x/tenthPow[rp])%10 {
			return false
		}

		rp--
		lp++
	}

	return true
}

func main() {

	fmt.Printf("121: %d\n", isPalindrome(131))
	fmt.Printf("-121: %d\n", isPalindrome(-141))
	fmt.Printf("10: %d\n", isPalindrome(10))
	fmt.Printf("1234554321: %d\n", isPalindrome(1234554321))

}
