// https://leetcode.com/problems/3sum/

/*
Explanation:

Solution for this problem is not that easy as it seems from
the first glance. We ahve to think about how this triplents
are made, this will guide us to the idea of sorting array and
using negative numbers as starting point. Then solution is pretty
basic using 2 pointers.

!!! Dont be shy to somethims use heave sorting, as in some cases !!!
!!! it can be more efficent then think of very complex algorithm !!!
!!! logic...                                                     !!!
*/

package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	resArr := [][]int{}

	numsLen := len(nums)
	negInd := 0

	for negInd < numsLen && nums[negInd] <= 0 {
		negNum := nums[negInd]

		lp := negInd + 1
		rp := numsLen - 1

		for lp < rp {
			curSum := negNum + nums[lp] + nums[rp]

			if curSum < 0 {
				lp++
			} else if curSum > 0 {
				rp--
			} else {
				resArr = append(resArr, []int{negNum, nums[lp], nums[rp]})

				lp++
				for nums[lp] == nums[lp-1] && lp < rp {
					lp++
				}
			}
		}

		negInd++
		for negInd < numsLen && nums[negInd] == nums[negInd-1] {
			negInd++
		}
	}

	return resArr
}

func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))

}
