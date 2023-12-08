package main

import (
	"fmt"

	"gorgonia.org/gorgonia"
	"gorgonia.org/tensor"
)

func main() {
	// Create a computation graph
	g := gorgonia.NewGraph()

	// Define tensors
	x := gorgonia.NewTensor(g, tensor.Float64, 2, gorgonia.WithShape(2, 2), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithBacking([]float64{1, 2, 3, 4}))))

	y := gorgonia.NewTensor(g, tensor.Float64, 2, gorgonia.WithShape(2, 2), gorgonia.WithValue(tensor.New(tensor.Of(tensor.Float64), tensor.WithBacking([]float64{5, 6, 7, 8}))))

	// Perform addition operation
	z, err := gorgonia.Add(x, y)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Compile graph
	machine := gorgonia.NewTapeMachine(g)

	// Execute graph
	if err := machine.RunAll(); err != nil {
		fmt.Println(err)
		return
	}

	// Get result
	result := z.Value().Data().([]float64)
	fmt.Println("Result of addition:", result)
}
