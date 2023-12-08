package main

import "fmt"

func returnHeighestElement(nums []int) int {
	maxValue := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
		}
	}

	return maxValue
}

func returnMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func returnArrSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	return sum
}

func trap(height []int) int {
	leftArr := []int{}
	rightArr := []int{}
	sumArr := []int{}
	leftArr = append(leftArr, height[0])
	for i := 0; i < len(height); i++ {
		if i != 0 {
			maxValues := returnHeighestElement(height[:i])
			if maxValues > height[i] {
				leftArr = append(leftArr, maxValues)
			} else {
				leftArr = append(leftArr, height[i])
			}
		}
		maxValuesForLeft := returnHeighestElement(height[i:])
		if maxValuesForLeft > height[i] {
			rightArr = append(rightArr, maxValuesForLeft)
		} else {
			rightArr = append(rightArr, height[i])
		}
		minValue := returnMin(leftArr[i], rightArr[i])
		minValue = minValue - height[i]
		sumArr = append(sumArr, minValue)
	}
	return returnArrSum(sumArr)
}

func main() {
	fmt.Println("Hello world")
	array := []int{3, 1, 2, 4, 0, 1, 3, 2}
	fmt.Println(trap(array))
}
