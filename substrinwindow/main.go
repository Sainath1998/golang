package main

import (
	"fmt"
)

func maxOfArr(nums []int) int {
	maxval := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxval {
			maxval = nums[i]
		}
	}
	return maxval
}

func maxSlidingWindow(nums []int, k int) []int {
	maxArr := []int{}
	left := 0
	right := k
	for right <= len(nums) {
		val := maxOfArr(nums[left:right])
		maxArr = append(maxArr, val)
		left++
		right++
	}
	return maxArr
}

func maxSlidingWindow2(nums []int, k int) []int {
	maxArr := []int{}
	i := 0
	for i < len(nums) {
		maxArr = append(maxArr, maxOfArr(nums[:k]))
		val := nums[0]
		nums = nums[1:]
		nums = append(nums, val)
	}
	return maxArr
}

func main() {
	s := []int{1, 3, -1, -3, 5, 3, 6, 7}
	t := 3
	result := maxSlidingWindow(s, t)
	fmt.Println(result)
}
