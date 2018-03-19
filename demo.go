package main

import (
	"flag"

	"github.com/gonum/matrix/mat64"
	goplotjs "github.com/rginestou/goplotjs/v1"
)

// Run the demo
// a flag "--demo=[line, matrix]" is required to select the plot type
func main() {
	demo := flag.String("demo", "line", "line, matrix, map")
	flag.Parse()

	switch *demo {
	case "line":
		plotLines()
	case "matrix":
		plotMatrix()
	}

	// Show the graph
	goplotjs.Show("window")
}

func plotLines() {
	// Define settings
	size := 100
	dx := 0.01

	// Allocate memory
	x := make([]float64, size)
	y1 := make([]float64, size)
	y2 := make([]float64, size)

	// Fill data
	for i := 0; i < size; i++ {
		x[i] = dx * float64(i)
		y1[i] = float64(i * i)
		y2[i] = 2 * float64(i*i)
	}

	// Call plot routines
	goplotjs.PlotLine(x, y1, "x²")
	goplotjs.AddPlotLine(x, y2, "2x²")
	goplotjs.SetTitle("Quadratic")
}

func plotMatrix() {
	// Define settings
	size := 100
	dx := 0.01
	dy := 0.05

	// Allocate memory
	m := mat64.NewDense(size, size, nil)
	x := mat64.NewVector(size, nil)
	y := mat64.NewVector(size, nil)
	// x_array := make([]float64, size)
	// y_array := make([]float64, size)
	for i := 0; i < size; i++ {
		x.SetVec(i, dx*float64(i))
		y.SetVec(i, dy*float64(i))
		// x_array[i] = dx * float64(i)
		// y_array[i] = dy * float64(i)
	}

	// Fill data
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m.Set(i, j, float64(i*j))
		}
	}

	// Call plot routines
	// goplotjs.PlotMatrix(x, y, m)
	// goplotjs.SetTitle("Heatmap Vector Axis")

	// goplotjs.PlotDense(m, x_array, y_array)
	// goplotjs.SetTitle("Heatmap Array Axis")

	// goplotjs.PlotDense(m, dx, dy, -0.5, 1.5)
	// goplotjs.SetTitle("Heatmap Increment-Offset Axis")

}

// 	case "map":
// 		// Allocate memory
// 		lat := []float64{4.413841, 4.423841}
// 		lon := []float64{50.860589, 50.862589}

// 		// Call plot routines
// 		goplotjs.PlotMap(lat, lon)
// 		goplotjs.SetTitle("Brussels Map")

// 		goplotjs.Show(true)
