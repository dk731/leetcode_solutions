// https://leetcode.com/problems/add-two-numbers/

/*
Explanation:

1. Copy l1, l2 values to local variables
2. Create result head and point local pointer to it
3. Iterate while at least one of the input pointer is not nil:
	1. Create new list node and move current point to it
	2. Determine and set new value to the current node
4. !!! In the end check overflow, if not null -> create new node !!!
*/

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resHead := &ListNode{0, nil}
	resP := resHead
	p1 := l1
	p2 := l2
	overflow := 0

	for p1 != nil || p2 != nil {
		resP.Next = &ListNode{0, nil}
		resP = resP.Next

		tempVal := overflow
		overflow = 0

		if p1 != nil {
			tempVal += p1.Val
			p1 = p1.Next
		}

		if p2 != nil {
			tempVal += p2.Val
			p2 = p2.Next
		}

		resP.Val = tempVal % 10
		if tempVal > 9 {
			overflow = 1
		}
	}

	if overflow > 0 {
		resP.Next = &ListNode{overflow, nil}
	}

	return resHead.Next
}

func main() {
	num1 := ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	num2 := ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	res := addTwoNumbers(&num1, &num2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}
