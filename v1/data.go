package goplotjs

// Graph main graph structure
type Graph struct {
	Title string        `json:"title"`
	Type  string        `json:"type"`
	Data  []interface{} `json:"data"`
}

// DataLine for LINE plot
type DataLine struct {
	Name string    `json:"name"`
	X    []float64 `json:"x"`
	Y    []float64 `json:"y"`
}

// DataMatrix for MATRIX plot (heatmap)
type DataMatrix struct {
	Name   string      `json:"name"`
	X      []float64   `json:"x"`
	Y      []float64   `json:"y"`
	Matrix [][]float64 `json:"matrix"`
}
