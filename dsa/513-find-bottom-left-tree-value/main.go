package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Queue represents a FIFO data structure.
type Queue struct {
	items []*TreeNode
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item *TreeNode) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
func (q *Queue) Dequeue() *TreeNode {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	return len(q.items)
}

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var node *TreeNode
	q := Queue{}
	q.Enqueue(root)

	for !q.IsEmpty() {
		node = q.Dequeue()
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
	}
	return node.Val
}

func main() {
	fmt.Println(findBottomLeftValue(&TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}))
}
