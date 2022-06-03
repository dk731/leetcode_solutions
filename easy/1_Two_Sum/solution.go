// https://leetcode.com/problems/two-sum/

/*
Explanation:

Create hash-map lookup table where [target - current]{index}
Iterate over array of numbers, check if current number exists
in lookup table:
	1. If yes -> we found solution
	2. If no -> add current number to the table
*/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {

		if val, ok := m[v]; ok {
			return []int{val, i}
		} else {
			m[target-v] = i
		}

	}

	return []int{}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
