// https://leetcode.com/problems/contains-duplicate/

/*
Explanation:

Simple solution using set for with O(n) time complexity.

!!! Solution of constant space is available sorting and !!!
!!! checking adjesent numbers                           !!!
*/

package main

import "fmt"

func containsDuplicate(nums []int) bool {
	checkMap := map[int]struct{}{}

	for _, val := range nums {
		if _, ok := checkMap[val]; ok {
			return true
		}

		checkMap[val] = struct{}{}
	}

	return false
}

func main() {

	fmt.Printf("[1,2,3,1]: %d\n", containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Printf("[1,2,3,4]: %d\n", containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Printf("[1,1,1,3,3,4,3,2,4,2]: %d\n", containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))

}
