// https://leetcode.com/problems/longest-consecutive-sequence/

/*
Explanation:

To solve in linear time, think of a way use additional memory, for
better speed. Use set of numbers, the iterate over input numbers,
and check that number is a sequency start (no left neighbour), then
iterate while number is inside set. Update max len

!!! Keep in mind that using hashmap as set makes lookup table      !!!
!!! creation time O(n)                                             !!!
*/

package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	lookupTable := map[int]struct{}{}
	maxLen := 0

	for _, val := range nums {
		lookupTable[val] = struct{}{}
	}

	for _, val := range nums {
		// Check if start
		if _, ok := lookupTable[val-1]; !ok {
			curLen := 0
			for i := val; ; i++ {
				if _, ok := lookupTable[i]; !ok {
					break
				}
				curLen++
			}

			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}

	return maxLen
}
func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
}
