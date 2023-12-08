package main

import "fmt"

func GetMin(nums []int) int {
	min := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > min {
			min = nums[i]
		}
	}

	return min

}

func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		fmt.Println(prices[i:])
		value := GetMin(prices[i:])
		profitValue := value - prices[i]
		if profitValue > profit {
			profit = profitValue
		}
	}
	return profit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}
