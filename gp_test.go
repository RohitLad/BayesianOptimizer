package main

import (
	"BayesianOptimizer/gaussianprocesses"
	"fmt"
	"math"
	"math/rand"
)

func f(x, y float64) float64 {
	return math.Cos(x/2)/2 + math.Sin(y/4)
}

func main() {
	gp := gaussianprocesses.Init()

	for i := 0; i < 200; i++ {
		a := rand.Float64()*2*math.Pi - math.Pi
		b := rand.Float64()*2*math.Pi - math.Pi
		ip := []float64{a, b}
		op := f(a, b)
		gp.InsertSingleInput(ip, op)
	}

	ip := [][]float64{{0.25, 0.75}, {0.5, 0.5}}
	mean, variance, _ := gp.Predict(ip)
	fmt.Println("Mean: ", mean, "  variance: ", variance)
	fmt.Println("feval: ", f(0.25, 0.75), "feval: ", f(0.5, 0.5))
}
