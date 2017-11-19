package main

import (
	"github.com/Romaingin/goplotjs/v1"
	"github.com/gonum/matrix/mat64"
	"flag"
)

func main()  {
	demo := flag.String("demo", "line", "line, matrix, map")
	flag.Parse()

	switch *demo {
	case "matrix":
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
		goplotjs.PlotDense(m, x, y)
		goplotjs.SetTitle("Heatmap & Contour")
		goplotjs.Show(true)

	case "line":
		// Define settings
		size := 100
		dx := 0.01

		// Allocate memory
		x1 := make([]float64, size)
		y1 := make([]float64, size)
		x2 := make([]float64, size)
		y2 := make([]float64, size)

		// Fill data
		for i := 0; i < size; i++ {
			x1[i] = dx * float64(i)
			y1[i] = float64(i * i)

			x2[i] = dx * float64(i)
			y2[i] = 2 * float64(i * i)
		}

		// Call plot routines
		goplotjs.PlotLine(x1, y1, "x²")
		goplotjs.AddPlotLine(x2, y2, "2x²")
		goplotjs.SetTitle("Quadratic")
		goplotjs.Show(true)

	case "map":
		// Allocate memory
		lat := []float64{4.413841, 4.423841}
		lon := []float64{50.860589, 50.862589}

		// Call plot routines
		goplotjs.PlotMap(lat, lon)
		goplotjs.SetTitle("Brussels Map")
		goplotjs.Show(true)
	case "multi":
		// Define settings
		size := 100
		dx := 0.01

		// Allocate memory
		x1 := make([]float64, size)
		y1 := make([]float64, size)
		x2 := make([]float64, size)
		y2 := make([]float64, size)

		// Fill data
		for i := 0; i < size; i++ {
			x1[i] = dx * float64(i)
			y1[i] = float64(i)

			x2[i] = dx * float64(i)
			y2[i] = float64(i * i)
		}

		// Call plot routines
		goplotjs.PlotLine(x1, y1, "x")
		goplotjs.SetTitle("Linear")

		goplotjs.PlotLine(x2, y2, "x²")
		goplotjs.SetTitle("Quadratic")
		goplotjs.Show(true)
	}
}
