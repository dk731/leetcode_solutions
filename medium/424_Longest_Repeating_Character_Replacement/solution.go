// https://leetcode.com/problems/longest-repeating-character-replacement/

/*
Explanation:

Silding window approach. On each itteration add current character
to the count table then find most frequent character in active widnow.
Check that windowLen - count(freqChar) is less than k, if yes this means
that current window is valid, expand window. If no then remove offset
character from count table and move whole window by increesing offset

!!! Remeber that technically O(n*26) ~ O(n)  !!!
*/

package main

import (
	"fmt"
)

func characterReplacement(s string, k int) int {
	windowLen := 0
	offset := 0
	countTable := make([]int, 26)
	sLen := len(s)

	for i := 0; i < 26; i++ {
		countTable[i] = 0
	}

	for offset+windowLen < sLen {
		currentChar := s[offset+windowLen]
		countTable[currentChar-65]++

		maxCount := -1
		maxChar := -1
		for i, count := range countTable {
			if count > maxCount {
				maxCount = count
				maxChar = i
			}
		}

		if windowLen-countTable[maxChar] < k {
			windowLen++
		} else {
			offsetChar := s[offset]
			offset++
			countTable[offsetChar-65]--
		}
	}

	return windowLen
}

func main() {
	fmt.Println("ABAB: ", characterReplacement("ABAB", 2))
	fmt.Println("AABABBA: ", characterReplacement("AABABBA", 1))
	fmt.Println("AAAA: ", characterReplacement("AAAA", 0))
}
