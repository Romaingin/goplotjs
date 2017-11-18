# GoPlotJS

_JS powered plot for GO_

Based on *gonum*.

```
go get gonum.org/v1/gonum/...
```

## Functions

**PlotDense(m _Dense_, x _[]float64_, y _[]float64_)** Plot a matrix of data whose axis are given by `x` and `y`. Dimensions must agree.

**PlotLine(x _[]float64_, y _[]float64_, name _string_)** Plot a function of `x` given `y`, with label `name`. Dimensions must agree.

**AddPlotLine(x _[]float64_, y _[]float64_, name _string_)** Add a function of `x` given `y`, with label `name` to the previous line plot. Dimensions must agree.

**PlotMap(lat _[]float64_, lon _[]float64_)** Plot a map using the *MapBox* API and superimpose points given their latitudes and longitudes. Dimensions must agree.


**SetTitle(title _string_)** Function to call after a new plot. Add a title to the graph.

**Show(block _bool_)** Function to call once at the end. Launch the visualization
accessible at `localhost:8080` in any browser.

## Usage

### Line plotting

```go
x1 := make([]float64, size)
y1 := make([]float64, size)
x2 := make([]float64, size)
y2 := make([]float64, size)
...

// Call plot routines
goplotjs.PlotLine(x1, y1, "x²")
goplotjs.AddPlotLine(x2, y2, "2x²")
goplotjs.SetTitle("Quadratic")
goplotjs.Show(true)
```

### Matrix plotting

```go
data := make([]float64, size*size)
...
m := mat64.NewDense(size, size, data)
x := make([]float64, size)
y := make([]float64, size)

goplotjs.PlotDense(m, x, y)
goplotjs.Show(true)
```
