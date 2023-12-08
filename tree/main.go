package main

import (
	"fmt"
)

type Node struct {
	Value       int
	left, right *Node
}

func (n *Node) insert(val int) *Node {
	if n == nil {
		n = &Node{Value: val, left: nil, right: nil}
		return n
	}
	if val < n.Value {
		n.left = n.left.insert(val)
	} else {
		n.right = n.right.insert(val)
	}
	return n
}

/* inorder traversal first left then current then right ALSO DFS*/
func (n *Node) displayInorder() {
	if n != nil {
		n.left.displayInorder()
		fmt.Println(n.Value)
		n.right.displayInorder()
	}
}

/* preorder traversal first current the left then right */

func (n *Node) displaypreorder() {
	if n != nil {
		fmt.Printf("%d \n", n.Value)
		n.left.displaypreorder()
		n.right.displaypreorder()
	}
}

/* POST order traversal firts left right then current */
func (n *Node) displayPostOrder() {
	if n != nil {
		n.left.displayPostOrder()
		n.right.displayPostOrder()
		fmt.Printf("%d \n", n.Value)
	}
}

/* bread first serach*/

func (n *Node) BFS(key int) bool {
	if n == nil {
		return false
	}
	queue := []*Node{n}
	fmt.Println(len(queue))
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Println(current.Value)

		if current.left != nil {
			return current.left.BFS(key)
		}

		if current.right != nil {
			return current.right.BFS(key)
		}
	}
	return false
}

func (n *Node) search(key int) bool {
	if n != nil {
		if n.Value == key {
			return true
		} else if key < n.Value {
			return n.left.search(key)
		} else {
			return n.right.search(key)
		}
	}
	return false
}

func (n *Node) depthFirstSearch() {
	if n != nil {
		fmt.Println(n.Value)

	}
}

func (n *Node) InvertBinarTree() {
	if n == nil {
		return
	}
	n.left, n.right = n.right, n.left
	n.left.InvertBinarTree()
	n.right.InvertBinarTree()
}

func (n *Node) calculateBreadthOfTree() {
	leftIndex := 0
	rightIndex := 0
	currentNode := n
	currentNode2 := n
	for currentNode.left != nil {
		currentNode = currentNode.left
		leftIndex++
	}

	for currentNode2.right != nil {
		currentNode2 = currentNode2.right
		rightIndex++
	}
}

func diameterOfBinaryTree(root *Node) int {
	// n := 0

	return -1
}

func (n *Node) rightView() {
	if n == nil {
		return
	}
	queue := []Node{*n}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[i]
			if i == size-1 {
				fmt.Println(curr.Value)
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
	if n == nil {
		return
	}
	queue := []Node{*n}
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[i]

			if i == 0 {
				fmt.Println(curr.Value)
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

func main() {
	a := Node{}
	a.insert(4)
	a.insert(7)
	a.insert(6)
	a.insert(2)
	a.insert(5)
	a.insert(1)
	a.insert(3)
	a.leftView()

}
