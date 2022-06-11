// https://leetcode.com/problems/3sum/

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
				for nums[lp] == nums[lp-1] && lp < numsLen {
					lp++
				}
			}
		}

		negInd++
		if negInd < numsLen && negInd > 0 && nums[negInd] == nums[negInd-1] {
			negInd++
		}
	}

	return resArr
}

func main() {

	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{}))

}
