package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var Dia int = 0

func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := list.New()
	queue.PushBack(root)
	i := 1

	for i < len(arr) {
		currentNode := queue.Remove(queue.Front()).(*TreeNode)

		// Left child
		if arr[i] != -200 { // Assuming -1 denotes None/-200 in Python code
			currentNode.Left = &TreeNode{Val: arr[i]}
			queue.PushBack(currentNode.Left)
		}
		i++

		// Right child
		if i < len(arr) && arr[i] != -200 {
			currentNode.Right = &TreeNode{Val: arr[i]}
			queue.PushBack(currentNode.Right)
		}
		i++
	}

	return root
}

func travel(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := travel(node.Left)
	right := travel(node.Right)
	Dia = max(Dia, left+right)

	return max(left, right) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	Dia = 0
	travel(root)
	return Dia
}

func main() {
	fmt.Println(diameterOfBinaryTree(buildTree([]int{1, 2})))
	fmt.Println(diameterOfBinaryTree(buildTree([]int{1, 2, 3, 4, 5})))

	p1 := buildTree([]int{4, -7, -3, -200, -200, -9, -3, 9, -7, -4, -200, 6, -200, -6, -6, -200, -200, 0, 6, 5, -200, 9, -200, -200, -1, -4, -200, -200, -200, -2})

	fmt.Println(diameterOfBinaryTree(p1))
}
