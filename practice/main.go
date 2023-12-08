package main

import (
	"fmt"
)

// ListNode represents a node in the linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// MergeTwoLists merges two sorted linked lists
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Create a dummy node to simplify the code
	dummy := &ListNode{}
	current := dummy

	// Continue until either of the lists becomes empty
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	// If there are remaining nodes in either list, append them to the result
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}

// DisplayList displays the elements of the linked list
func DisplayList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func returnMaxOfArr(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func minEatingSpeed(piles []int, h int) int {

	return -1
}

func main() {
	// // Create two sorted linked lists
	// list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	// list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: nil}}}

	// // Display the original lists
	// fmt.Print("List 1: ")
	// DisplayList(list1)
	// fmt.Print("List 2: ")
	// DisplayList(list2)

	// // Merge the lists
	// mergedList := MergeTwoLists(list1, list2)

	// // Display the merged list
	// fmt.Print("Merged List: ")
	// DisplayList(mergedList)
	piles := []int{1000000000, 1000000000}
	h := 3
	// minEatingSpeed(piles, h)
	fmt.Println(minEatingSpeed(piles, h))

}
