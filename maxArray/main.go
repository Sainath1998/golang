package main

import (
	"fmt"
)

func getLowestNumber(num1, num2 int) int {
	if num1 == num2 {
		return num1
	} else if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func maxArea(nums []int) int {
	left := 0
	right := len(nums) - 1
	maxAreaValue := 0
	for left < right {
		lh := nums[left]
		rh := nums[right]
		lowestNumber := getLowestNumber(lh, rh)
		area := lowestNumber * (right - left)
		if area > maxAreaValue {
			maxAreaValue = area
		}

		if lh > rh {
			for left != right && nums[right] <= rh {
				right--
			}
		} else {
			for left != right && nums[left] <= lh {
				left++
			}
		}

	}

	return maxAreaValue
}

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	fmt.Println(maxArea(a))
}
