// https://leetcode.com/problems/valid-parentheses/

/*
Explanation:

Simple solution using lookup table to fast map open/close
brackets. Create occured brackets stack, on each step check
if this stack is valid, if not return false immdediatly.

!!! Be cautious with edgde cases, make sure to check res stack !!!
!!! in the very end, if not empty, then there are left         !!!
!!! unclosed brackets
*/

package main

import "fmt"

var openBrackets = map[rune]rune{'(': ')', '[': ']', '{': '}'}
var closeBrackets = map[rune]rune{')': '(', ']': '[', '}': '{'}

func isValid(s string) bool {

	resStack := []rune{}

	for _, char := range s {
		if val, ok := openBrackets[char]; ok {
			resStack = append(resStack, val)
		} else {
			resLen := len(resStack)
			if resLen == 0 || resStack[resLen-1] != char {
				return false
			}

			// Pop last bracket
			resStack = resStack[0 : len(resStack)-1]
		}
	}

	return len(resStack) == 0
}

func main() {

	fmt.Printf("'()': %d\n", isValid("()"))
	fmt.Printf("'()[]{}': %d\n", isValid("()[]{}"))
	fmt.Printf("'(]': %d\n", isValid("(]"))

}
