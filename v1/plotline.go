package goplotjs

type VisualizeLine struct{
	Type string `json:"type"`
	Title string `json:"title"`
	X [][]float64 `json:"x"`
	Y [][]float64 `json:"y"`
	Names []string `json:"names"`
}

var lastVizualizationLine *VisualizeLine

// > PlotLine forms a visualizable data structures given some x ant y points.
func PlotLine(x, y []float64, name string) {
	if len(x) != len(y) {
		panic("Dimensions must agree")
	}

	var viz VisualizeLine
	viz.X = make([][]float64, 1)
	viz.Y = make([][]float64, 1)
	viz.Names = make([]string, 1)

	viz.Type = "line"
	viz.Title = "Line Plot"
	viz.X[0] = x
	viz.Y[0] = y
	viz.Names[0] = name

	lastVizualizationLine = &viz
	addVisualization(&viz)
}

// > AddPlotLine draw one function plot on top of the previous line chart
func AddPlotLine(x, y []float64, name string) {
	if len(x) != len(y) {
		panic("Dimensions must agree")
	}
	if lastVizualizationLine == nil {
		panic("No line plot to add to")
	}

	(*lastVizualizationLine).X = append((*lastVizualizationLine).X, x)
	(*lastVizualizationLine).Y = append((*lastVizualizationLine).Y, y)
	(*lastVizualizationLine).Names = append((*lastVizualizationLine).Names, name)
}

// Visualization interface getters and setters
func (v *VisualizeLine) setTitle(title string) {
	v.Title = title
}

func (v *VisualizeLine) getTitle() string {
	return v.Title
}

func (v *VisualizeLine) getType() string {
	return v.Type
}
