# GoPlotJS

_JS powered plot for GO_

Documentation : https://rginestou.github.io/goplotjs/

## Example

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
