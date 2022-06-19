// https://leetcode.com/problems/valid-parentheses/

/*
Explanation:

Solution consists of two parts, first to filter all exesive characters,
second: using simple 2 pointers approach, iterate from beggining and
end, on each step check if equal, if not return false

!!! For filtering use simple range check for character              !!!
*/

package main

import "fmt"

func isPalindrome(s string) bool {

	clearS := ""

	for _, v := range s {
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') && (v < '0' || v > '9') {
			continue
		}

		if v >= 'A' && v <= 'Z' {
			clearS += string(v + 32)
		} else {
			clearS += string(v)
		}
	}

	lp := 0
	rp := len(clearS) - 1

	fmt.Print(clearS)

	for lp < rp {
		if clearS[lp] != clearS[rp] {
			return false
		}

		lp++
		rp--
	}

	return true
}

func main() {

	fmt.Printf("'A man, a plan, a canal: Panama': %d\n", isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Printf("'race a car': %d\n", isPalindrome("race a car"))
	fmt.Printf("'0P': %d\n", isPalindrome("0P"))

}
