package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	result := []int{}
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return result
}

func main() {
	result := []int{2, 7, 11, 15}
	fmt.Println(twoSum(result, 9))
}
