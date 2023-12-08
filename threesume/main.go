package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		target := -nums[i]

		for left < right {
			if nums[left]+nums[right] == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, -2, 1, 3, -1, -4}
	fmt.Println(threeSum(nums))
}
