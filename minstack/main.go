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

func main() {
	minstack := Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	// minstack.Pop()
	// minstack.Top()
	fmt.Println(len(minstack.stack))

}
