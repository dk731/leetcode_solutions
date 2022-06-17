// https://leetcode.com/problems/product-of-array-except-self/

/*
Explanation:

To solve must use clever idea of prefixes and postfixes.
Hardes is to think about how excatcly to store information
about computed multiplication. For creating an O(1) space
complexity one more small twek must be made, that is storing
computed results in answer array.

!!! Do not forget that solution of O(n + n + n + n + n)     !!!
!!! reduces to the O(n), and it is valid to iterate over    !!!
!!! input array multiple times                              !!!
*/

package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	numsLen := len(nums)
	resArr := make([]int, numsLen)
	tempVal := 1

	for i := 0; i < numsLen; i++ {
		resArr[i] = tempVal
		tempVal *= nums[i]
	}

	tempVal = 1

	for i := numsLen - 1; i >= 0; i-- {
		resArr[i] *= tempVal
		tempVal *= nums[i]
	}

	return resArr
}

func main() {

	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))

}
