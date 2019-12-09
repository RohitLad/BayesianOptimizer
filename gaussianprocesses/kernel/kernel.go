package kernel

import (
	"gonum.org/v1/gonum/mat"
)

type ExpCov struct {
}

func (ec *ExpCov) CalcSymCov(X *mat.Dense) {

	N := X.Dims[0]
	C := mat.NewSymDense(N, nil)
	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {

			C.SetSym(i, j, v)
		}
	}

}
