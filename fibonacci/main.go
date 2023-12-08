package main

import "fmt"

// fibonacci of a number using dynanmic problem

func fibonacci(num int) []int {
	f := []int{0, 1}
	for i := 2; i <= num; i++ {
		f = append(f, f[i-1]+f[i-2])
	}
	return f
}

func main() {
	fmt.Print("Hello world")
	// fiboacci usingdp
	fmt.Println(fibonacci(10))
}
