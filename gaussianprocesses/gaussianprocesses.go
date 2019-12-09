package gaussianprocesses

import "BayesianOptimizer/gaussianprocesses/kernel"

type GP struct {
	inputs  [][]float64
	outputs []float64
	kernel  *kernel.ExpCov
	reset   bool
}

func (gp *GP) InsertSingleInput(input []float64, output float64) {
	defer func() {
		gp.reset = true
	}()
	gp.inputs = append(gp.inputs, input)
	gp.outputs = append(gp.outputs, output)
}

func (gp *GP) InsertBulkInputs(inputs [][]float64, outputs []float64) {

	N := len(inputs)
	for i := 0; i < N; i++ {
		gp.InsertSingleInput(inputs[i], outputs[i])
	}
}
