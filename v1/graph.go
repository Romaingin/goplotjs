package goplotjs

// Graphs store the different graphs to be plotted
type Graphs []Graph

func (gs *Graphs) addGraph(g Graph) {
	*gs = append(*gs, g)
}

func (gs *Graphs) setTitle(t string) {
	(*gs)[gs.len()-1].Title = t
}

func (gs Graphs) len() int {
	return len(gs)
}

func (gs *Graphs) get(i int) *Graph {
	return &(*gs)[i]
}

func (gs *Graphs) getLast() *Graph {
	return gs.get(gs.len() - 1)
}
