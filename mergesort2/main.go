package main

import "fmt"

func merger(left, right []int) []int {
	fmt.Println("*********")
	fmt.Println(left)
	fmt.Println(right)
	fmt.Println("*********")
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

	result = append(result, right[rightIndex:]...)
	result = append(result, left[leftIndex:]...)

	return result

}

func mergeSort(nums []int) []int {

	// Base condition if length of array less that 2 then return array itself
	if len(nums) <= 1 {
		return nums
	}

	// calculate the mid
	mid := len(nums) / 2

	// calculate the right
	right := nums[mid:]

	// calculate the left
	left := nums[:mid]

	// recursicvelsey dived left
	left = mergeSort(left)

	// recusivwely dived right
	right = mergeSort(right)

	// finally merger both left and right
	return merger(left, right)
}

func main() {
	a := []int{54, 2, 4, 1, 278, 1238, 123, 11}

	mergeSort(a)
}
