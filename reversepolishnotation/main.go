package main

import (
	"fmt"
	"math"
	"strconv"
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

func evaluate(nums []int, operator string) int {
	if operator == "*" {
		return nums[1] * nums[0]
	} else if operator == "+" {
		return nums[1] + nums[0]
	} else if operator == "-" {
		return nums[1] - nums[0]
	} else {
		return nums[1] / nums[0]

	}
}

func evalRPN(tokens []string) int {
	minstack := Constructor()
	for i := 0; i < len(tokens); i++ {
		ele := tokens[i]
		if ele != "*" && ele != "+" && ele != "-" && ele != "/" {
			converted, _ := strconv.Atoi(ele)
			minstack.Push(converted)

		} else {
			operands := []int{}
			topEle := minstack.Top()
			operands = append(operands, topEle)
			minstack.Pop()
			topEle2 := minstack.Top()
			operands = append(operands, topEle2)
			minstack.Pop()
			result := evaluate(operands, ele)
			minstack.Push(result)

		}
	}
	return minstack.Top()
}

func main() {
	stringsArr := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(stringsArr))
}
