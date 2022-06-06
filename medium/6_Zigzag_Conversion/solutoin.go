// https://leetcode.com/problems/zigzag-conversion/

/*
Explanation:

No magic stuff, to solve rewrite input example strings
output indexes, find logical approach to compute each index,
implement algorithm in code.

!!! Visualize, look at different approaches !!!
*/

package main

import "fmt"

func convert(s string, numRows int) string {
	sLen := len(s)
	resStr := make([]byte, sLen)
	firstStep := 0
	secondStep := 0
	oInd := 0
	iInd := 0

	for iRow := 0; iRow < numRows; iRow++ {
		firstStep = numRows*2 - 2
		secondStep = firstStep
		isFirst := true

		if iRow != 0 && iRow != numRows-1 {
			firstStep -= iRow * 2
			secondStep = (numRows*2 - 2) - firstStep
		}

		if firstStep <= 0 {
			firstStep = 1
			secondStep = 1
		}

		for iInd = iRow; iInd < sLen; {
			resStr[oInd] = s[iInd]
			oInd++

			if isFirst {
				iInd += firstStep
			} else {
				iInd += secondStep
			}

			isFirst = !isFirst
		}
	}

	return string(resStr)
}

func main() {

	fmt.Printf("PAYPALISHIRING; 3: %s\n", convert("PAYPALISHIRING", 3))
	fmt.Printf("PAYPALISHIRING; 4: %s\n", convert("PAYPALISHIRING", 4))
	fmt.Printf("A; 1: %s\n", convert("A", 1))

}
