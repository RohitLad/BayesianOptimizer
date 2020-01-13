package bo

import (
	"BayesianOptimizer/gaussianprocesses"

	"gonum.org/v1/gonum/optimize"
)

type BayesianOptimizer struct {
	Objective   func([]float64) float64
	verbose     int
	randomState int
	vm          *VariableManager
	gp          *gaussianprocesses.GP
	dims        int
	changed     bool
}

func New(Objective func([]float64) float64) *BayesianOptimizer {
	return &BayesianOptimizer{vm: NewVariableManager(), gp: gaussianprocesses.Init(), Objective: Objective}
}

func (bo *BayesianOptimizer) InsertNewVariable(Name string, Min, Max float64, IsInt bool) {
	bo.changed = true
	bo.vm.InsertNewVariable(Name, Min, Max, IsInt)
}

func (bo *BayesianOptimizer) Maximize(init_points, n_iter int, x0 []float64, util UtilityFunction, f func([]float64) float64) error {
	ok, err := bo.vm.VerifyBounds(x0)
	if !ok {
		return err
	}
	initFeval := f(x0)
	bo.vm.SaveMapping(x0, initFeval)
	for n := 0; n < n_iter; n++ {
		x, y, err := bo.vm.PeepMapping(n)
		if err != nil {
			return err
		}
		bo.gp.InsertSingleInput(x, y)

	}

}

func (bo BayesianOptimizer) nextParameter(random bool, util UtilityFunction, evaluatedLoss, nRestarts int) ([]float64, error) {

	minAcquisition := 1.0
	bestSol := []float64{0.}

	for i := 0; i < nRestarts; i++ {
		startPoint := bo.vm.RamdomSample()
		problem := optimize.Problem{
			Func: func(x []float64) float64 {
				return util.Estimate(x, 0., 0.) // Fill in here
			},
		}
		result, err := optimize.Minimize(problem, startPoint, nil, &optimize.LBFGS{})
		if err != nil {
			return []float64{0.}, err
		}
		if result.F < minAcquisition {
			minAcquisition = result.F
			bestSol = result.X
		}
	}
	return bestSol, nil
}
