// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

/*
Explanation:

Classic sliding window problem. Create 2 pointers (left and right), iterate
while right pointer is in bound, on each step calculate new profit and
update max profit. Move appropriate pointer based on profit sign. In the
end check if pointers are equal increese right pointer.

!!! Solution fro first glance seems counter-intuitive and illogical !!!
!!! seems like left pointer sometimes might not be updated, but     !!!
!!! iterating alghorithm it comes clear that biggest dip will find  !!!
!!! will find its self                                              !!!
*/

package main

import "fmt"

func maxProfit(prices []int) int {
	lp := 0
	rp := 1
	maxProfit := 0
	pricesLen := len(prices)

	for rp < pricesLen {
		tmpProfit := prices[rp] - prices[lp]
		if tmpProfit > maxProfit {
			maxProfit = tmpProfit
		}

		if tmpProfit < 0 {
			lp++
		} else {
			rp++
		}

		if rp == lp {
			rp++
		}
	}

	return maxProfit
}

func main() {

	fmt.Printf("[7,1,5,3,6,4]: %d\n", maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Printf("[7,6,4,3,1]: %d\n", maxProfit([]int{7, 6, 4, 3, 1}))

}
