package bo

import (
	"BayesianOptimizer/gaussianprocesses"
	"math"
)

// x -> point at which utility function is evaluated
// ymax -> maximum eval yet
// xi -> Exploration/Exploitation tradeoff param

type UtilityFunction interface {
	Estimate(x []float64, ymax, xi float64) float64
}

var _ UtilityFunction = UCB{}

type UCB struct {
	gp    *gaussianprocesses.GP
	kappa float64
}

func (ucb UCB) Estimate(x []float64, ymax, xi float64) float64 {
	ip := [][]float64{x}
	mean, std, _ := ucb.gp.Predict(ip)
	return mean.At(0, 0) + std.At(0, 0)*ucb.kappa
}

var _ UtilityFunction = EI{}

var UnitNormal = Normal{Mu: 0, Sigma: 1.0}

type EI struct {
	gp *gaussianprocesses.GP
}

func (ei EI) Estimate(x []float64, ymax, xi float64) float64 {
	ip := [][]float64{x}
	mean, std, _ := ei.gp.Predict(ip)
	num := mean.At(0, 0) - ymax - xi
	z := (num) / math.Max(std.At(0, 0), 1e-8)
	return num*UnitNormal.CDF(z) + std.At(0, 0)*UnitNormal.PDF(z)
}

var _ UtilityFunction = POI{}

type POI struct {
	gp *gaussianprocesses.GP
}

func (poi POI) Estimate(x []float64, ymax, xi float64) float64 {
	ip := [][]float64{x}
	mean, std, _ := poi.gp.Predict(ip)
	z := (mean.At(0, 0) - ymax - xi) / math.Max(std.At(0, 0), 1e-8)
	return UnitNormal.CDF(z)
}
