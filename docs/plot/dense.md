---
title: Dense
---

## PlotDense

(m _Dense_, x _[]float64_, y _[]float64_)

Plot a heat-map corresponding to the matrix of data.
The axis `x` and `y` can be given as `mat64.Vector`, `[]float64`, or as increment with offset.

```go
var x, y mat64.Vector
goplotjs.PlotDense(m, x, y)
```

```go
var x, y []float64
goplotjs.PlotDense(m, x, y)
```

```go
var dx, dy float64
var offset_x, offset_dy float64
goplotjs.PlotDense(m, dx, dy, offset_x, offset_y)
```
