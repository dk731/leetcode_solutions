// https://leetcode.com/problems/median-of-two-sorted-arrays/
package main

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	p1 := 0
	p2 := 0
	s1 := len(nums1)
	s2 := len(nums2)

	curInd := 0
	startInd := (s1 + s2) / 2
	stopInd := startInd

	resNum := 0

	// Check if prime
	if (s1+s2)&1 == 0 {
		startInd--
	}

	for curInd <= stopInd && (p1 < s1 || p2 < s2) {
		curNum := 0
		if p1 < s1 && (p2 >= s2 || nums1[p1] < nums2[p2]) {
			curNum = nums1[p1]
			p1++
		} else {
			curNum = nums2[p2]
			p2++
		}

		if curInd == startInd {
			resNum += curNum
		}

		if curInd == stopInd {
			resNum += curNum
		}

		curInd++
	}

	return float64(resNum) / 2
}

func main() {

	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))

}
