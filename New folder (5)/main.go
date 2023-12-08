package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(node *TreeNode) {
	// Base case: check if the current node is nil
	if node == nil {
		return
	}

	// Process the current node (e.g., print its value)
	fmt.Println(node.Val)

	// Recursively visit the left and right subtrees
	dfs(node.Left)
	dfs(node.Right)
}

func main() {
	// Example binary tree
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
	}

	// Perform depth-first search
	fmt.Println("Depth-First Search:")
	dfs(root)
}
