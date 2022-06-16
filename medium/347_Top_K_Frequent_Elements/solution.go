// https://leetcode.com/problems/top-k-frequent-elements/

/*
Explanation:

Interesting extention to the 'counting' problems. At first glance
solution is straightforward and very basic: create countTable and
sort it in decending order, tak first k elements, but it is far from
the O(n) time complexity. To avoid unnescassery actions, we utilize
a frequency map (key is min occurences in input, value is array of
this values). Iterating over input we are updating out frequency map,
and keeping  track of most frequnt and at the same time suitable array.
In the end return this suitable array

!!! If cannot think of how to improve time complexity by just improving !!!
!!! an algorith, think about how to utilize additional memory to        !!!
!!! shortcut some parts of logic                                        !!!
*/

package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	resultsTable := map[int][]int{}
	countTable := map[int]int{}
	suitableFreq := 0

	for _, n := range nums {
		countTable[n]++
		curCount := countTable[n]

		if _, ok := resultsTable[curCount]; ok {
			resultsTable[curCount] = append(resultsTable[curCount], n)
		} else {
			resultsTable[curCount] = []int{n}
		}

		if curCount > suitableFreq && len(resultsTable[curCount]) >= k {
			suitableFreq = curCount
		}
	}

	return resultsTable[suitableFreq]
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(topKFrequent([]int{3, 3, 3, 0, 1, 0, 2}, 1))
	fmt.Println(topKFrequent([]int{1}, 1))
}
