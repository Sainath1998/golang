package main

import "fmt"

func merge(left, right []int) []int {
	leftIndex, righIndex := 0, 0
	result := []int{}

	for leftIndex < len(left) && righIndex < len(right) {
		if left[leftIndex] < right[righIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[righIndex])
			righIndex++
		}
	}

	result = append(result, left[leftIndex:]...)
	result = append(result, right[righIndex:]...)

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
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := mergeSort(a)
	fmt.Println(result)
}
