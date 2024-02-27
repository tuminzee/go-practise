package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to check if two binary trees are the same.
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// If both trees are empty, they are the same.
	if p == nil && q == nil {
		return true
	}
	// If one of the trees is empty but not the other, they are not the same.
	if p == nil || q == nil {
		return false
	}
	// If the current node values are not equal, return false.
	if p.Val != q.Val {
		return false
	}
	// Recursively check left and right subtrees.
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	// Example 1
	p1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Println("Example 1:", isSameTree(p1, q1)) // Output: true

	// Example 2
	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	fmt.Println("Example 2:", isSameTree(p2, q2)) // Output: false

	// Example 3
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 1}}
	q3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	fmt.Println("Example 3:", isSameTree(p3, q3)) // Output: false
}
