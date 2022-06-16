// https://leetcode.com/problems/group-anagrams/

/*
Explanation:

Extending 49 solution, keep track of each word countTable, which
is used to store word in resulting hash map, where key is countTable
and value is array of strings.

!!! Same as in 49 solution, instead of storing countTable, sorting !!!
!!! can be used. Have additional question about input data, as it  !!!
!!! is able to optimise solution                                   !!!
*/

package main

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	resTable := map[[26]int][]string{}
	countTable := [26]int{}

	for _, word := range strs {
		for i := 0; i < 26; i++ {
			countTable[i] = 0
		}

		for _, c := range word {
			tmp := c - 'a'
			countTable[tmp]++
		}

		if _, ok := resTable[countTable]; ok {
			resTable[countTable] = append(resTable[countTable], word)
		} else {
			resTable[countTable] = []string{word}
		}
	}

	finalArray := [][]string{}

	for _, resArr := range resTable {
		finalArray = append(finalArray, resArr)
	}

	return finalArray
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
