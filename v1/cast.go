package goplotjs

import "github.com/gonum/matrix/mat64"

func castToFloat64(x interface{}) []float64 {
	var X []float64

	switch v := x.(type) {
	case []float64:
		X = v
	case *mat64.Vector:
		X = v.RawVector().Data
	default:
		panic("Invalid parameter x")
	}

	return X
}

func castToMatrix64(m interface{}) [][]float64 {
	var M [][]float64

	switch v := m.(type) {
	case [][]float64:
		M = v
	case *mat64.Dense:
		r, _ := v.Dims()

		M := make([][]float64, r)
		for i := 0; i < r; i++ {
			M[i] = v.RawRowView(i)
		}
	default:
		panic("Invalid parameter matrix")
	}

	return M
}
