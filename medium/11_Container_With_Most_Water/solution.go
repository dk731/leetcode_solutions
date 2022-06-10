// https://leetcode.com/problems/container-with-most-water/

/*
Explanation:

Two pointers approach, start from outer side of an array.
On each step calculate and update max area, then move that pointer
which is limiting current area (that which is smaller)

!!! No need to overthink, solution is much simpler than it seems to be !!!
*/

package main

import (
	"fmt"
)

func maxArea(height []int) int {
	maxRes := -1
	lp := 0
	rp := len(height) - 1

	for lp < rp {
		curArea := (rp - lp)
		if height[lp] < height[rp] {
			curArea *= height[lp]
		} else {
			curArea *= height[rp]
		}

		if curArea > maxRes {
			maxRes = curArea
		}

		if height[lp] < height[rp] {
			lp++
		} else {
			rp--
		}
	}

	return maxRes
}

func main() {

	fmt.Printf("%d\n", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Printf("%d\n", maxArea([]int{1, 1}))

}
