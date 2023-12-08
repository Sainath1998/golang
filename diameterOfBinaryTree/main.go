package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 20
	fmt.Println(int(math.Max(float64(a), float64(b))))
}
