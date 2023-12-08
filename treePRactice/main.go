package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	Val         int
	left, right *Node
}

func (n *Node) insert(num int) *Node {
	if n == nil {
		n = &Node{Val: num, left: nil, right: nil}
		return n
	}
	if n.Val > num {
		n.left = n.left.insert(num)
	} else {
		n.right = n.right.insert(num)
	}
	return n
}

// Inorder traversal
func (n *Node) inorderTraversal() {
	if n != nil {
		n.left.inorderTraversal()
		fmt.Println(n.Val)
		n.right.inorderTraversal()
	}
	return
}

// preorder traversal
func (n *Node) preorderTraversal() {
	if n != nil {
		fmt.Println(n.Val)
		n.left.preorderTraversal()
		n.right.preorderTraversal()
	}
	return
}

// postorder traversal
func (n *Node) postordertraversal() {
	if n != nil {
		n.left.postordertraversal()
		n.right.postordertraversal()
		fmt.Println(n.Val)
	}
	return
}

// Level order traversal
func (n *Node) levelorderTraversal() {
	if n == nil {
		return
	}
	queue := []Node{*n}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			fmt.Println(curr.Val)
			if curr.left != nil {
				queue = append(queue, *curr.left)
			}

			if curr.right != nil {
				queue = append(queue, *curr.right)
			}
		}
		queue = queue[size:]

	}
}

// Right view of the binary tree
func (n *Node) RightView() {
	queue := []Node{*n}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			if i == size-1 {
				fmt.Println(curr.Val)
			}
			if curr.left != nil {
				queue = append(queue, *curr.left)
			}

			if curr.right != nil {
				queue = append(queue, *curr.right)
			}
		}
		queue = queue[size:]
	}
}

func (n *Node) leftView() {
	queue := []Node{*n}
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[i]

			if i == 0 {
				fmt.Println(curr.Val)
			}

			if curr.left != nil {
				queue = append(queue, *curr.left)
			}

			if curr.right != nil {
				queue = append(queue, *curr.right)
			}
		}

		queue = queue[size:]
	}
}

func maxPathSum(root *Node) int {
	var calculate func(root *Node) int
	maxSum := math.MinInt64
	calculate = func(root *Node) int {
		if root == nil {
			return 0
		}
		leftSum := int(math.Max(float64(calculate(root.left)), 0))
		rightSum := int(math.Max(float64(calculate(root.right)), 0))
		if root.Val+leftSum+rightSum > maxSum {
			maxSum = int(math.Max(float64(maxSum), float64(root.Val+leftSum+rightSum)))
		}
		return root.Val + int(math.Max(float64(leftSum), float64(rightSum)))
	}
	calculate(root)
	return maxSum
}

func max(i, sum int) {
	panic("unimplemented")
}

func (m *Node) levelorderTraversal2() string {
	result := []int{}
	queue := []Node{*m}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[i]
			result = append(result, curr.Val)

			if curr.left != nil {
				queue = append(queue, *curr.left)
			}

			if curr.right != nil {
				queue = append(queue, *curr.right)
			}
		}
		queue = queue[size:]
	}
	var stringArray []string
	for _, num := range result {
		stringArray = append(stringArray, strconv.Itoa(num))
	}
	newVal := strings.Join(stringArray, ", ")
	return newVal
}

func main() {
	var m *Node
	m = m.insert(1)
	m = m.insert(2)
	m = m.insert(3)
	m.levelorderTraversal2()
	// fmt.Println("inorder traversal")
	// m.inorderTraversal()
	// fmt.Println("Preortder traversal")
	// m.preorderTraversal()
	// fmt.Println("postordertraversal")
	// m.postordertraversal()
	// fmt.Println("levelorderTraversal")
	// m.levelorderTraversal()
	// fmt.Println("Left view")
	// m.leftView()
	// fmt.Println("Right view")
	// m.RightView()

}
