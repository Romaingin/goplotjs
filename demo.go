package main

import (
	"github.com/Romaingin/goplotjs/v1"
	"github.com/gonum/matrix/mat64"
	"fmt"
)

func main()  {
	// Define settings
	size := 100
	dx := 0.01
	dy := 0.05

	// Allocate memory
	data := make([]float64, size*size)
	x := make([]float64, size)
	y := make([]float64, size)

	// Fill data
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			data[size * i + j] = float64(i*j)
		}

		x[i] = dx * float64(i)
		y[i] = dy * float64(i)
	}
	m := mat64.NewDense(size, size, data)

	// Call plot routines
	fmt.Println("Matrix Visualization")
	goplotjs.PlotDense(m, x, y)
	goplotjs.Show(true)
}
