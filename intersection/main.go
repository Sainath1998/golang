package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	largest := nums1
	smallest := nums2
	if len(nums1) < len(nums2) {
		largest = nums2
		smallest = nums1
	}
	largestMap := make(map[int]int)

	for i := 0; i < len(largest); i++ {
		largestMap[largest[i]]++
	}
	result := []int{}
	for i := 0; i < len(smallest); i++ {
		if largestMap[smallest[i]] > 0 {
			result = append(result, smallest[i])
			largestMap[smallest[i]]--
		}
	}
	return result
}

func findTheDifference(s string, t string) byte {
	if len(s) <= 0 {
		return t[len(t)-1]
	}
	charmap := make(map[rune]bool)

	for _, val := range s {
		charmap[val] = true
	}

	for _, val2 := range t {
		if !charmap[val2] {
			return byte(val2)
		}
	}
	return 'a'
}

func merge(left, right []int) []int {
	leftIndex, rightIndex := 0, 0
	result := []int{}
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
	right := nums[mid:]
	left := nums[:mid]

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left, right)
}

func sortColors(nums []int) {
	nums = mergeSort(nums)
	fmt.Println(nums)

}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	sortColors(a)
}
