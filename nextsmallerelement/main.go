package main

import "fmt"

func nextSmallerElement(arr []int) []int {
	n := len(arr)
	result := make([]int, n)
	stack := make([]int, 0)

	for i := 0; i < n; i++ {
		for len(stack) > 0 && arr[i] < arr[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[top] = arr[i]
		}
		stack = append(stack, i)
	}
	for _, index := range stack {
		result[index] = -1
	}

	return result
}

// Abive function willl rteytn you the smakker elements in a stack

func main() {
	intArr := []int{2, 1, 4, 3}
	fmt.Println(nextSmallerElement(intArr))
}
