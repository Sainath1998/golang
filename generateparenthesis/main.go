package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	positiveInfinity := math.MaxInt
	for _, value := range this.stack {
		if value < positiveInfinity {
			positiveInfinity = value
		}
	}
	return positiveInfinity
}

func dailyTemperatures(temperatures []int) []int {
	dayArr := []int{}
	for i := 0; i < len(temperatures); i++ {
		currecntTemp := temperatures[i]
		nextTempValues := temperatures[i:]
		days := 0
		for j := 0; j < len(nextTempValues); j++ {
			if nextTempValues[j] > currecntTemp {
				days = j
				break
			}
		}
		dayArr = append(dayArr, days)

	}
	return dayArr

}

func main() {
	a := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(a))

}
