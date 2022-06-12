// https://leetcode.com/problems/remove-nth-node-from-end-of-list/submissions/

/*
Explanation:

Very simple 2 pointers approach, first init our pointers, then move
front pointer to needed offset, and iterate both pointers until front
one is null


!!! Be careful with pointer initialization, must think more about !!!
!!! what is happening in empty/1 element long lists               !!!
*/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head}
	p1 := newHead
	p2 := head

	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	for ; p2 != nil; p2 = p2.Next {
		p1 = p1.Next
		println(p1.Val)
	}

	p1.Next = p1.Next.Next

	return newHead.Next
}

func main() {
	testList := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	resList := removeNthFromEnd(&testList, 2)

	fmt.Print("[1,2,3,4,5], n = 2 -> ")
	for curp := resList; curp != nil; curp = curp.Next {
		fmt.Printf("%d; ", curp.Val)
	}
}
