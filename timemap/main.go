package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
	result = append(result, nums2[rightIndex:]...)
	result = append(result, nums1[leftIndex:]...)
	fmt.Println(result)
	length := len(result)
	mid := len(result) / 2
	if length%2 == 0 {
		medianVal := float64(result[mid]+result[mid-1]) / 2.0
		fmt.Println(medianVal)
		return medianVal
	}
	return float64(result[mid])
}
func main() {
	a := []int{1, 3}
	b := []int{2, 7}
	fmt.Println(findMedianSortedArrays(a, b))
}
