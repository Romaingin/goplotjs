package main

import (
	"github.com/Romaingin/goplotjs/v1"
	"github.com/gonum/matrix/mat64"
	m "math"
)

func gauss(x float64) float64 {
	sigma2 := 0.01
	return 0.1 * m.Exp(-(x-0.5)*(x-0.5) / sigma2)
}

func main() {
	// Simulation settings
	l_interval := 80
	t_interval := 100
	// uf := 10.0
	// w := 0.5
	dx := 1.0 / float64(l_interval)
	dt := 1.0 / float64(t_interval)
	x_0 := 0.0

	// Allocate memory
	t_ := mat64.NewVector(t_interval+1, nil)
	x_ := mat64.NewVector(l_interval+1, nil)
	for n := 0; n < t_interval+1; n++ {
		t_.SetVec(n, dt * float64(n))
	}
	for j := 0; j < l_interval+1; j++ {
		x_.SetVec(j, dx * float64(j) + x_0)
	}
	k := mat64.NewDense(t_interval+1, l_interval+1, nil)

	// Boundary conditions
	for j := 0; j < l_interval; j++ {
		k.Set(0, j, gauss(x_.At(j, 0)))
	}

	// Run simulation
	for n := 0; n < t_interval; n++ {
		c := 0.8
		for j := 1; j < l_interval; j++ {
			k.Set(n+1,j, k.At(n,j) + dt * c*c * (k.At(n,j+1) - 2*k.At(n,j) + k.At(n,j-1)) * dt / (2*dx*dx) - c * (k.At(n,j+1) - k.At(n,j-1)) * dt / (2*dx))
		}
		k.Set(n+1,0, k.At(n,0) + dt * c*c * (k.At(n,1) - 2*k.At(n,0) + k.At(n,l_interval-1)) * dt / (2*dx*dx) - c * (k.At(n,1) - k.At(n,l_interval-1)) * dt / (2*dx))
		k.Set(n+1,l_interval, k.At(n+1,0))
	}

	// Call plot routines
	goplotjs.PlotDense(k, x_, t_)
	goplotjs.SetTitle("Wave Propagation")
	goplotjs.Show(true)
}
