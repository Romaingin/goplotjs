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
func PlotDense(m *mat64.Dense, x, y []float64) {
	r, c := m.Dims()
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

func (v *VisualizeDense) setTitle(title string) {
	v.Title = title
}
