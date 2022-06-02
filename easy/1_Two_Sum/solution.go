// https://leetcode.com/problems/two-sum/
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
