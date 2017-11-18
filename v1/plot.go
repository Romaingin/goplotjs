package goplotjs

type Viz interface {
	setTitle(string)
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
