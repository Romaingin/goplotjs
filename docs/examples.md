---
title: Examples
---

### Line Plot

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

### Heat-map Plot

```go
m := mat64.NewDense(size, size, data)
x := make([]float64, size)
y := make([]float64, size)
...

// Call plot routines
goplotjs.PlotDense(m, x, y)
goplotjs.Show(true)
```
