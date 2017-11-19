package goplotjs

// Define interface for all types of graphs
type Viz interface {
	setTitle(string)
	getTitle() string
	getType() string
}

var (
	visualizations []Viz
	lastVizualizationItem *Viz
)

func addVisualization(v Viz)  {
	lastVizualizationItem = &v
	visualizations = append(visualizations, v)
}

func SetTitle(title string) {
	if lastVizualizationItem == nil {
		panic("Title set before plot")
	}

	(*lastVizualizationItem).setTitle(title)
}
