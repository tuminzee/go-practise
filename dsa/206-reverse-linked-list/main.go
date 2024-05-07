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

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	printLinkedList(head)
	printLinkedList(reverseList(head))

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
