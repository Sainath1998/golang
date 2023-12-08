package main

import (
	"fmt"
	"sort"
)

type MinStack struct {
	stack []float32
}

func Constructor() MinStack {
	return MinStack{
		stack: make([]float32, 0),
	}
}

func (this *MinStack) Push(val float32) {
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() float32 {
	return this.stack[len(this.stack)-1]
}

func carFleet(target int, position []int, speed []int) int {
	arr := [][]int{}

	for i := 0; i < len(position); i++ {
		speedDistArr := []int{}
		speedDistArr = append(speedDistArr, position[i])
		speedDistArr = append(speedDistArr, speed[i])
		arr = append(arr, speedDistArr)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] > arr[j][0]
	})
	minstack := Constructor()

	for _, element := range arr {
		time := float32(target-element[0]) / float32(element[1])
		minstack.Push(time)
		if len(minstack.stack) >= 2 && minstack.stack[len(minstack.stack)-1] <= minstack.stack[len(minstack.stack)-2] {
			minstack.Pop()
		}
	}
	fmt.Println(minstack.stack)
	return len(minstack.stack)
}

func main() {
	// minstack := Constructor()
	// minstack.Push(-2)
	// minstack.Push(0)
	// minstack.Push(-3)
	// // minstack.Pop()
	// // minstack.Top()
	// fmt.Println(minstack.stack)
	// fmt.Println(minstack.stack[len(minstack.stack)-2])
	target := 10
	position := []int{6, 8}
	speed := []int{3, 2}
	fmt.Println(carFleet(target, position, speed))

}
