package goplotjs

// PlotMatrix set the graph to MATRIX mode, and add matrix data
// on the fly
func PlotMatrix(m, x, y interface{}, offset ...float64) {
	M := castToMatrix64(m)

	// Dimensions
	r := len(M)
	c := len(M[0])

	var X, Y []float64

	// Special cast in case of offset
	if vx, ok := x.(float64); ok {
		if len(offset) != 2 {
			panic("Invalid offset")
		}
		X = make([]float64, c)
		for n := 0; n < c; n++ {
			X[n] = vx*float64(n) + offset[0]
		}
	} else {
		X = castToFloat64(x)
	}

	if vy, ok := y.(float64); ok {
		if len(offset) != 2 {
			panic("Invalid offset")
		}
		X = make([]float64, c)
		for n := 0; n < c; n++ {
			Y[n] = vy*float64(n) + offset[0]
		}
	} else {
		Y = castToFloat64(y)
	}

	// Check dimension
	if len(X) != c {
		panic("Dimensions must agree")
	}
	if len(Y) != r {
		panic("Dimensions must agree")
	}

	// Set graph
	var g Graph
	g.Type = "matrix"
	g.Data = make([]interface{}, 1)

	g.Data[0] = DataMatrix{
		X:      X,
		Y:      Y,
		Matrix: M,
	}

	graphs.addGraph(g)
}
