package main

import (
	m "math"

	"github.com/gonum/matrix/mat64"
	"github.com/rginestou/goplotjs/v1"
)

func gauss(x float64) float64 {
	sigma2 := 0.01
	return 0.1 * m.Exp(-(x-0.5)*(x-0.5)/sigma2)
}

func main() {
	// Simulation settings
	lInterval := 80
	tInterval := 100
	// uf := 10.0
	// w := 0.5
	dx := 1.0 / float64(lInterval)
	dt := 1.0 / float64(tInterval)
	x0 := 0.0

	// Allocate memory
	t := mat64.NewVector(tInterval+1, nil)
	x := mat64.NewVector(lInterval+1, nil)
	for n := 0; n < tInterval+1; n++ {
		t.SetVec(n, dt*float64(n))
	}
	for j := 0; j < lInterval+1; j++ {
		x.SetVec(j, dx*float64(j)+x0)
	}
	k := mat64.NewDense(tInterval+1, lInterval+1, nil)

	// Boundary conditions
	for j := 0; j < lInterval; j++ {
		k.Set(0, j, gauss(x.At(j, 0)))
	}

	// Run simulation
	for n := 0; n < tInterval; n++ {
		c := 0.8
		for j := 1; j < lInterval; j++ {
			k.Set(n+1, j, k.At(n, j)+dt*c*c*(k.At(n, j+1)-2*k.At(n, j)+k.At(n, j-1))*dt/(2*dx*dx)-c*(k.At(n, j+1)-k.At(n, j-1))*dt/(2*dx))
		}
		k.Set(n+1, 0, k.At(n, 0)+dt*c*c*(k.At(n, 1)-2*k.At(n, 0)+k.At(n, lInterval-1))*dt/(2*dx*dx)-c*(k.At(n, 1)-k.At(n, lInterval-1))*dt/(2*dx))
		k.Set(n+1, lInterval, k.At(n+1, 0))
	}

	// Call plot routines
	goplotjs.PlotMatrix(k, x, t)
	goplotjs.SetTitle("Wave Propagation")
	goplotjs.Show()
}
