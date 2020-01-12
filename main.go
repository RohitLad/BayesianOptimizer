package main

import (
	"BayesianOptimizer/bo"
	"fmt"
	"math"
)

func BlackBox(x []float64) float64 {
	return -math.Pow(x[0], 2) - math.Pow(x[1]-1, 2) + 1
}

func main() {
	optimizer := bo.New(BlackBox)
	optimizer.InsertNewVariable("X", 2, 4, false)
	optimizer.InsertNewVariable("Y", -3, 3, false)
	//optimizer.Maximize(2, 3)
	fmt.Println(test([]float64{0.0, 1.0}, bo.UCB{}))
}

func test(x []float64, fc bo.UtilityFunction) float64 {
	return 0.
}
