package goplotjs

import (
	"github.com/gonum/matrix/mat64"
)

type VisualizeDense struct{
	Type string `json:"type"`
	Title string `json:"title"`
	Data [][]float64 `json:"data"`
	X []float64 `json:"x"`
	Y []float64 `json:"y"`
}

func DenseToFloat64(m *mat64.Dense) [][]float64 {
	r, _ := m.Dims()

	t := make([][]float64, r)
	for i := 0; i < r; i++ {
		t[i] = m.RawRowView(i)
	}

	return t
}

// > PlotDense forms a visualizable data structures based on
// a matrix, given some x ant y axis.
// x and y can be vectors, []floats or increments
func PlotDense(m *mat64.Dense, x_, y_ interface{}, offset... float64) {
	r, c := m.Dims()
	var x, y []float64

	// Cast x
	switch v := x_.(type) {
	case []float64:
		x = v
	case *mat64.Vector:
		x = v.RawVector().Data
	case float64:
		if len(offset) != 2 {
			panic("Invalid offset")
		}
		x = make([]float64, c)
		for n := 0; n < c; n++ {
			x[n] = v * float64(n) + offset[0]
		}
	default:
		panic("Invalid parameter x")
	}

	// Cast y
	switch v := y_.(type) {
	case []float64:
		y = v
	case *mat64.Vector:
		y = v.RawVector().Data
	case float64:
		if len(offset) != 2 {
			panic("Invalid offset")
		}
		y = make([]float64, r)
		for n := 0; n < r; n++ {
			y[n] = v * float64(n) + offset[1]
		}
	default:
		panic("Invalid parameter x")
	}

	if len(x) != c {
		panic("Dimensions must agree")
	}
	if len(y) != r {
		panic("Dimensions must agree")
	}

	var viz VisualizeDense
	viz.Type = "matrix"
	viz.Title = "Matrix Plot"
	viz.Data = DenseToFloat64(m)
	viz.X = x
	viz.Y = y

	addVisualization(&viz)
}

// Visualization interface getters and setters
func (v *VisualizeDense) setTitle(title string) {
	v.Title = title
}

func (v *VisualizeDense) getTitle() string {
	return v.Title
}

func (v *VisualizeDense) getType() string {
	return v.Type
}
