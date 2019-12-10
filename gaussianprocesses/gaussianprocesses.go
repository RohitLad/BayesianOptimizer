package gaussianprocesses

import (
	"BayesianOptimizer/gaussianprocesses/kernel"
	"fmt"

	"github.com/pkg/errors"
	"gonum.org/v1/gonum/mat"
)

type GP struct {
	inputs  [][]float64
	outputs []float64
	kernel  kernel.ExpCov
	Kinv    mat.Dense
	CholK   mat.Cholesky // technically not required
	reset   bool
}

func Init() *GP {
	gp := GP{kernel: kernel.ExpCov{SigmaF: 1.0, L: 1.0, SigmaN: 0.0}, reset: true}
	return &gp
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

func (gp *GP) Train() error {
	K := gp.kernel.CalcSymCov(gp.inputs)
	if ok := gp.CholK.Factorize(K); !ok {
		return errors.Wrap(errors.New("Failed to Factorize"), "Train")
	}
	I := mat.NewSymDense(len(gp.inputs), nil)
	for i := 0; i < len(gp.inputs); i++ {
		I.SetSym(i, i, 1.0)
	}
	gp.Kinv = *mat.NewDense(len(gp.inputs), len(gp.inputs), nil)
	gp.CholK.SolveTo(&gp.Kinv, I)
	gp.reset = false
	return nil
}

func (gp *GP) Predict(inputs [][]float64) (*mat.Dense, *mat.SymDense, error) {
	if gp.reset {
		err := gp.Train()
		fmt.Println("This is the error:: ", err)
	}
	Kss := gp.kernel.CalcSymCov(inputs)
	Kst := gp.kernel.CalcASymCov(gp.inputs, inputs)
	ymat := mat.NewDense(len(gp.inputs), 1, gp.outputs)
	Npredict := len(inputs)
	N := len(gp.inputs)
	KstKyinv := mat.NewDense(Npredict, N, nil)
	KstKyinv.Mul(Kst.T(), &gp.Kinv)
	muPredict := mat.NewDense(Npredict, 1, nil)
	muPredict.Mul(KstKyinv, ymat)
	sigPredict := mat.NewSymDense(Npredict, nil)
	ak := mat.NewSymDense(Npredict, nil)
	for i := 0; i < Npredict; i++ {
		for j := i; j < Npredict; j++ {
			v := -1 * mat.Dot(KstKyinv.RowView(i), Kst.ColView(j))
			ak.SetSym(i, j, v)
		}
	}
	sigPredict.AddSym(Kss, ak)

	return muPredict, sigPredict, nil
}
