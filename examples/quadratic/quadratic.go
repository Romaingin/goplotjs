package main

import (
	goplotjs "github.com/rginestou/goplotjs/v1"
)

func main() {
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

	goplotjs.PlotLine(x, y1, "x²")
	goplotjs.SetTitle("Quadratic 2")

	// Show the graph
	goplotjs.Show("window")
}

// 	case "map":
// 		// Allocate memory
// 		lat := []float64{4.413841, 4.423841}
// 		lon := []float64{50.860589, 50.862589}

// 		// Call plot routines
// 		goplotjs.PlotMap(lat, lon)
// 		goplotjs.SetTitle("Brussels Map")

// 		goplotjs.Show(true)
