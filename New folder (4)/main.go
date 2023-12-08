package main

import "fmt"

func merge(left, right []int) []int {
	result := []int{}
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)
	return result
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := nums[:mid]
	right := nums[mid:]
	left = mergeSort(left)
	right = mergeSort(right)
	return merge(left, right)
}

func main() {
	a := []int{41, 24, 12, 441, 12, 5621, 12, 999}
	fmt.Print(mergeSort(a))
}
