package bo

import (
	"BayesianOptimizer/gaussianprocesses"
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
