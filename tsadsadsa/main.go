package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	result := []int{}
	leftIndex, rightIndex := 0, 0

	for leftIndex < len(nums1) && rightIndex < len(nums2) {
		if nums1[leftIndex] < nums2[rightIndex] {
			result = append(result, nums1[leftIndex])
			leftIndex++
		} else {
			result = append(result, nums2[rightIndex])
			rightIndex++
		}
	}
	result = append(result, nums1[leftIndex:]...)
	result = append(result, nums2[rightIndex:]...)
	nums1 = result
	fmt.Println(nums1)
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
}
