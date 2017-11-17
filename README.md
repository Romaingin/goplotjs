# GoPlotJS

_JS powered plot for GO_

Based on *gonum*.

```
go get gonum.org/v1/gonum/...
```

## Functions

**PlotDense(m _Dense_, x _[]float64_, y _[]float64_)** Plot a matrix of data whose axis are given by `x` and `y`. Dimensions must agree.

**Show(block _bool_)** Function to call once at the end. Launch the visualization
accessible at `localhost:8080` in any browser.

## Usage

### Matrix plotting

```go
data := make([]float64, size*size)
m := mat64.NewDense(size, size, data)
x := make([]float64, size)
y := make([]float64, size)

goplotjs.PlotDense(m, x, y)
goplotjs.Show(true)
```
