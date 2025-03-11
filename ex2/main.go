package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{Val: 0}
	curr := dummyHead
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyHead.Next
}

func main() {
	// Create first number: 342 (represented as 2 -> 4 -> 3)
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// Create second number: 465 (represented as 5 -> 6 -> 4)
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	// Add the two numbers
	result := addTwoNumbers(l1, l2)

	// Print the result
	for result != nil {
		fmt.Print(result.Val)
		if result.Next != nil {
			fmt.Print(" -> ")
		}
		result = result.Next
	}
}
