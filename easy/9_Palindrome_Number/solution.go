// https://leetcode.com/problems/palindrome-number/

/*
Explanation:

A better approach is available by creating inverted number,
and returning if it is equal to the original (with this approach
even bigger shortcut can be made, by reversing only half of the
digits). This solution is taking approach very simillar to the
5. longest palindromic substring, by comparing extracting numbers
one by one and comparing them

!!! If better approach for 10'th power calculation could be found !!!
!!! in come cases this solution would be better that soring whole !!!
!!! number                                                        !!!
*/

package main

import "fmt"

func tenthPow(x int) int {
	resVal := 1
	for i := 0; i < x; i++ {
		resVal *= 10
	}

	return resVal
}

func isPalindrome(x int) bool {
	// Exclude nagative numbers from the very beggining
	if x < 0 {
		return false
	}

	digitsCount := 1

	for x/tenthPow(digitsCount) >= 1 {
		digitsCount++
	}

	rp := digitsCount - 1
	lp := 0

	for lp < rp {
		rpRes := tenthPow(rp)
		lpRes := tenthPow(lp)

		// fmt.Printf("Res: %d->%d   %d->%d\n", (x/rpRes)%10, rpRes, (x%lpRes*10)/(lpRes), lpRes)
		if (x/rpRes)%10 != (x/lpRes)%10 {
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
