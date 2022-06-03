// https://leetcode.com/problems/longest-substring-without-repeating-characters/

/*
Explanation:

Use slightly modified version of sliding window.
Whenever it is possible, expand window by adding new
values from the right side. If not possible, shrink
window by incrementing starting possition (remove from the left).

!!! On each window expansion, register maximum value for result !!!

For better performance use array of int, with length 255 -> ASCII lookup
*/

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	var currentWindow [255]int

	for i := 0; i < 255; i++ {
		currentWindow[i] = 0
	}

	currentOffset := 0
	windowLength := 0
	strLength := len(s)

	maxLength := 0

	for currentOffset+windowLength < strLength {
		curChar := s[currentOffset+windowLength]

		if currentWindow[curChar] == 0 {
			currentWindow[curChar] = 1
			windowLength++

			if windowLength > maxLength {
				maxLength = windowLength
			}
		} else {
			currentWindow[s[currentOffset]]--
			currentOffset++
			windowLength--
		}
	}

	return maxLength
}

func main() {

	fmt.Printf("abcabcbb: %d\n", lengthOfLongestSubstring("abcabcbb"))
	fmt.Printf("bbbbb: %d\n", lengthOfLongestSubstring("bbbbb"))
	fmt.Printf("pwwkew: %d\n", lengthOfLongestSubstring("pwwkew"))

}
