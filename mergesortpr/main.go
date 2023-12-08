// Merge sort in golang

package main

import "fmt"

func merge(left, right []int) []int {
	leftIndex, rightIndex := 0, 0
	resultArr := []int{}
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			resultArr = append(resultArr, left[leftIndex])
			leftIndex++
		} else {
			resultArr = append(resultArr, right[rightIndex])
			rightIndex++
		}
	}
	// Attach remaiing elements to result arr
	resultArr = append(resultArr, left[leftIndex:]...)
	resultArr = append(resultArr, right[rightIndex:]...)
	return resultArr
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
	fmt.Println("Hello world")
	a := []int{1, 5, 2, 6, 1, 2, 78, 21, 53, 12}
	result := mergeSort(a)
	fmt.Println(result)
}

// 1 1 2 2 5 6 12 21 53 78] this is a sorted arrary and merge sort provide nlog(n) time complexity in all casses
