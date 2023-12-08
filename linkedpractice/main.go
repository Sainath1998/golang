package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (li *ListNode) insert(num int) *ListNode {
	data := &ListNode{Val: num, Next: nil}
	if li == nil {
		return data
	}
	curr := li
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = data
	return li
}

func (li *ListNode) display() {
	curr := li
	for curr != nil {
		fmt.Printf("%d -->", curr.Val)
		curr = curr.Next
	}
}

func (li *ListNode) swap() *ListNode {
	curr := li
	for curr != nil && curr.Next != nil {
		curr.Val, curr.Next.Val = curr.Next.Val, curr.Val
		curr = curr.Next.Next
	}
	return li
}

func (li *ListNode) delete(num int) *ListNode {
	if num < 0 {
		return li
	}
	dummy := &ListNode{Next: li}
	for li.Next.Next != nil {
		li = li.Next
	}
	val := li.Next.Val
	li.Next = nil
	dummy2 := &ListNode{Val: val, Next: dummy.Next}
	return dummy2.delete(num - 1)

}

func rotate(li *ListNode, num int) *ListNode {
	if num == 0 {
		return li
	}
	dummy := &ListNode{Next: li}
	for li.Next.Next != nil {
		li = li.Next
	}
	val := li.Next.Val
	li.Next = nil
	dummy2 := &ListNode{Val: val, Next: dummy.Next}
	return rotate(dummy2, num-1)

}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	} else if head.Next == nil {
		return head
	}
	lenght := 1
	c := head
	for c.Next != nil {
		c = c.Next
		lenght++
	}
	if lenght == k {
		return head
	}
	newK := k % lenght
	return rotate(head, newK)
}

func nextPermutation(nums []int) [][]int {
	var backTrack func(nums []int, index int)
	resultArr := [][]int{}
	backTrack = func(nums []int, index int) {
		if len(nums) == index {
			temp := make([]int, len(nums))
			copy(temp, nums)
			resultArr = append(resultArr, temp)
		}
		for i := index; i < len(nums); i++ {
			nums[i], nums[index] = nums[index], nums[i]
			backTrack(nums, index+1)
			nums[index], nums[i] = nums[i], nums[index]
		}

	}
	backTrack(nums, 0)
	return resultArr
}

func ArrSum(arr []int) int {
	sum := arr[0]
	for i := 1; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	var bactrack func(nums []int, index int)

	bactrack = func(candidates []int, index int) {
		if ArrSum(candidates) == target {
			result = append(result, candidates)
		} else {
			for i := 0; i < len(candidates); i++ {
				bactrack(candidates[i:], i)
				bactrack(candidates[i+1:], i+1)
			}
		}
	}
	return result
}

func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1
	leftArr := []int{}
	righArr := []int{}
	for left < len(nums) && right >= 0 {
		if nums[left] == target && len(leftArr) == 0 {
			leftArr = append(leftArr, left)
			left++
		} else {
			left++
		}
		if nums[right] == target && len(righArr) == 0 {
			righArr = append(righArr, right)
			right--
		} else {
			right--
		}
	}
	fmt.Println(leftArr)
	// fmt.Println(righArr)
	newArr := []int{}
	newArr = append(newArr, leftArr...)
	newArr = append(newArr, righArr...)
	if len(newArr) > 0 {
		return newArr
	}
	return []int{-1, -1}
}

func arrNum(nums []int, key int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == key {
			return true
		}
	}
	return false
}

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	largest := nums[len(nums)-1]
	if largest < 0 {
		return 1
	}
	for i := 1; i < largest; i++ {
		if !arrNum(nums, i) {
			return i
		}
	}

	return nums[len(nums)-1] + 1
}

func main() {
	// fmt.Println("Hello world")

	// var a *ListNode
	// a = a.insert(1)
	// a = a.insert(2)
	// a = a.insert(3)
	// a = a.insert(4)
	// a = a.insert(5)
	// a.display()
	// fmt.Println("")

	// lenght := 1
	// c := a
	// for c.Next != nil {
	// 	c = c.Next
	// 	lenght++
	// }
	// fmt.Println(lenght)
	// a = a.delete(20000)
	// // a = rotateRight(a, 1)
	// a.display()
	nums := []int{-5}
	// target := 1
	result := firstMissingPositive(nums)
	fmt.Println(result)
}
