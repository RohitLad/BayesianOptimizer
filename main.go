package main

import (
	"BayesianOptimizer/gaussianprocesses"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	ak := gaussianprocesses.GP{}
	v := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	A := mat.NewDense(3, 3, v)
	fmt.Println(A)
	fa := mat.Formatted(A, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}
