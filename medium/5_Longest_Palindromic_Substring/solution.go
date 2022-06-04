// https://leetcode.com/problems/longest-palindromic-substring/

/*
Explanation:

Straightforward solution of iterating over input string,
each iteration id is the center of the current palidromic
substring, expand as long as it posible. Check if current
substring has lenth bigger than resulting length, update.

!!! Handle odd and event substrings exactly the same, except !!!
!!! starting idxes                                           !!!
*/

package main

import "fmt"

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	}

	curStart := 0

	bestLength := 1
	bestStart := 0

	for curStart <= sLen {
		tmpL := curStart
		tmpR := curStart

		for tmpL >= 0 && tmpR < sLen && s[tmpL] == s[tmpR] {
			tmpL--
			tmpR++
		}

		if tmpR-tmpL > bestLength {
			bestStart = tmpL
			bestLength = tmpR - tmpL
		}

		tmpL = curStart
		tmpR = curStart + 1

		for tmpL >= 0 && tmpR < sLen && s[tmpL] == s[tmpR] {
			tmpL--
			tmpR++
		}

		if tmpR-tmpL > bestLength {
			bestStart = tmpL
			bestLength = tmpR - tmpL
		}

		curStart++
	}

	return s[bestStart+1 : bestStart+bestLength]
}

func main() {

	fmt.Printf("babad: %s\n", longestPalindrome("babad"))
	fmt.Printf("cbbd: %s\n", longestPalindrome("cbbd"))
	fmt.Printf("ccc: %s\n", longestPalindrome("ccc"))

}
