package goplotjs

type VisualizeMap struct{
	Type string `json:"type"`
	Title string `json:"title"`
	Lat []float64 `json:"lat"`
	Lon []float64 `json:"lon"`
}

// > PlotMap
func PlotMap(lat, lon []float64) {
	if len(lat) != len(lon) {
		panic("Dimensions must agree")
	}

	var viz VisualizeMap
	viz.Type = "map"
	viz.Title = "Matrix Plot"
	viz.Lat = lat
	viz.Lon = lon

	addVisualization(&viz)
}

func (v *VisualizeMap) setTitle(title string) {
	v.Title = title
}

func (v *VisualizeMap) getTitle() string {
	return v.Title
}

func (v *VisualizeMap) getType() string {
	return v.Type
}
