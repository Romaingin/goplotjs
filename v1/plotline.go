package goplotjs

// PlotLine set the graph to LINE mode, and add line data
// on the fly
func PlotLine(x, y interface{}, name string) {
	X := castToFloat64(x)
	Y := castToFloat64(y)

	if len(X) != len(Y) {
		panic("Dimensions must agree")
	}

	// Set graph
	var g Graph
	g.Type = "line"
	g.Data = make([]interface{}, 1)

	g.Data[0] = DataLine{
		Name: name,
		X:    X,
		Y:    Y,
	}

	graphs.addGraph(g)
}

// AddPlotLine ...
func AddPlotLine(x, y interface{}, name string) {
	X := castToFloat64(x)
	Y := castToFloat64(y)

	if len(X) != len(Y) {
		panic("Dimensions must agree")
	}

	g := graphs.getLast()

	g.Data = append(g.Data, DataLine{
		Name: name,
		X:    X,
		Y:    Y,
	})
}
