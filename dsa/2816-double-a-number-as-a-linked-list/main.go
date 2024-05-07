package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkedList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("None")
}

func reverseList(head *ListNode) *ListNode {
	current := head
	var back *ListNode

	for current != nil {
		front := current.Next
		current.Next = back
		back = current
		current = front
	}
	return back
}

func doubleIt(head *ListNode) *ListNode {
	rev := reverseList(head)
	dummy := rev

	digit, carry := 0, 0
	for rev != nil {
		digit = rev.Val*2 + carry
		val := digit % 10
		carry = digit / 10

		rev.Val = val
		if rev.Next == nil && carry > 0 {
			rev.Next = &ListNode{
				Val:  carry,
				Next: nil,
			}
			break
		}
		rev = rev.Next

	}

	return reverseList(dummy)
}

func main() {
	head := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}

	printLinkedList(head)
	printLinkedList(doubleIt(head))

}
