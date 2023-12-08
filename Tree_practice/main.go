package main

import "fmt"

type Node struct {
	Val         int
	left, right *Node
}

func (n *Node) insert(nums int) *Node {
	if n == nil {
		n = &Node{Val: nums, left: nil, right: nil}
		return n
	} else if n.Val < nums {
		n.right = n.right.insert(nums)
	} else {
		n.left = n.left.insert(nums)
	}
	return n
}

func (n *Node) inorder() {
	if n != nil {
		n.left.inorder()
		fmt.Println(n.Val)
		n.right.inorder()
	}
	return
}

func (n *Node) preorder() {
	if n != nil {
		fmt.Println(n.Val)
		n.left.preorder()
		n.right.preorder()
	}
	return
}

func (n *Node) postorder() {
	if n != nil {
		n.left.postorder()
		n.right.postorder()
		fmt.Println(n.Val)
	}
	return
}

func (n *Node) levelordertraversal() {
	queue := []*Node{n}
	result := [][]int{}
	for len(queue) > 0 {
		size := len(queue)
		arr := []int{}
		for i := 0; i < size; i++ {
			curr := queue[i]
			arr = append(arr, curr.Val)
			if curr.left != nil {
				queue = append(queue, curr.left)
			}

			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
		queue = queue[size:]
		result = append(result, arr)
	}

	fmt.Println(result)

}

func (n *Node) leftView() {
	queue := []*Node{n}
	result := [][]int{}
	for len(queue) > 0 {
		size := len(queue)
		arr := []int{}
		for i := 0; i < size; i++ {
			curr := queue[i]
			if i == 0 {
				arr = append(arr, curr.Val)
			}

			if curr.left != nil {
				queue = append(queue, curr.left)
			}

			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
		queue = queue[size:]
		result = append(result, arr)
	}

	fmt.Println(result)

}

func (n *Node) rightView() {
	queue := []*Node{n}
	result := [][]int{}
	for len(queue) > 0 {
		size := len(queue)
		arr := []int{}
		for i := 0; i < size; i++ {
			curr := queue[i]
			if i == size-1 {
				arr = append(arr, curr.Val)
			}

			if curr.left != nil {
				queue = append(queue, curr.left)
			}

			if curr.right != nil {
				queue = append(queue, curr.right)
			}
		}
		queue = queue[size:]
		result = append(result, arr)
	}

	fmt.Println(result)

}

func inorderTraversal(root *Node) []int {
	result := []int{}
	var inorder func(root *Node)

	inorder = func(root *Node) {
		if root != nil {
			inorder(root.left)
			result = append(result, root.Val)
			inorder(root.right)
		}
	}

	inorder(root)
	return result
}

func (n *Node) insertv(nums int) *Node {
	if n == nil {
		n = &Node{Val: nums, left: nil, right: nil}
		return n
	} else if n.Val < nums {
		n.left = n.left.insert(nums)
	} else {
		n.right = n.right.insert(nums)
	}
	return n
}

// func sortedArrayToBST(nums []int) *Node {
// 	var tree *Node

// 	for i := 0; i < len(nums); i++ {
// 		tree = tree.insert(nums[i])
// 	}
// 	return tree
// }

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal2(root *TreeNode) []int {
	result := []int{}
	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode) {
		if root != nil {
			inorder(root.Left)
			result = append(result, root.Val)
			inorder(root.Left)
		}
	}

	inorder(root)
	return result
}

func sortedArrayToBST(nums []int) *TreeNode {
	var tree *TreeNode
	var insert func(tree *TreeNode, num int) *TreeNode
	insert = func(tree *TreeNode, num int) *TreeNode {
		if tree == nil {
			return &TreeNode{Val: num, Left: nil, Right: nil}
		} else if tree.Val < num {
			tree.Left = insert(tree.Left, num)
		} else {
			tree.Right = insert(tree.Right, num)
		}
		return tree
	}
	for i := 0; i < len(nums); i++ {
		insert(tree, nums[i])
	}
	return tree
}
func main() {
	var tree *Node
	tree = tree.insert(10)

	x := inorderTraversal(tree)
	fmt.Println(x)
}
