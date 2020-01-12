package bo

import "BayesianOptimizer/gaussianprocesses"

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

func (bo *BayesianOptimizer) Maximize(init_points, n_iter int, util UtilityFunction) {
	for i := 0; i < init_points; i++ {
		sample := bo.vm.RamdomSample()
		output := bo.Objective(sample)
		bo.gp.InsertSingleInput(sample, output)
	}
}
