package goplotjs

import (
	"github.com/gonum/matrix/mat64"
)

type VisualizeData struct{
	Type string `json:"matrix"`
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
func PlotDense(m *mat64.Dense, x, y []float64) {
	var viz VisualizeData
	viz.Type = "matrix"
	viz.Title = "Matrix Plot"
	viz.Data = DenseToFloat64(m)
	viz.X = x
	viz.Y = y

	visualizations = append(visualizations, viz)
}
