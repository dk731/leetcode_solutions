// https://leetcode.com/problems/valid-anagram/submissions/

/*
Explanation:

1. O(1) memory:
    Sort both words, compare strings

2. O(n) time:
    Create lookup table, where store each letter of both words
in the end check that this table has all zeros

!!! Simillar to the 217, we have 2 choices of contant memory !!!
!!! or O(n) speed.                                           !!!
*/

package main

import "fmt"

func isAnagram(s string, t string) bool {
	charsArr := [255]byte{}
	for i := 0; i < 255; i++ {
		charsArr[i] = 0
	}

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		charsArr[s[i]]++
		charsArr[t[i]]--
	}

	for _, val := range charsArr {
		if val != 0 {
			return false
		}
	}

	return true
}

func main() {

	fmt.Printf("'anagram'/'nagaram': %d\n", isAnagram("anagram", "nagaram"))
	fmt.Printf("'rat'/'car': %d\n", isAnagram("rat", "car"))
	fmt.Printf("'QWEqweewq'/'QqWwEeqwe': %d\n", isAnagram("QWEqweewq", "QqWwEeqwe"))

}
