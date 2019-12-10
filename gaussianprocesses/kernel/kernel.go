package kernel

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type ExpCov struct {
	SigmaF float64
	L      float64
	SigmaN float64
}

func (ec *ExpCov) CalcSymCov(X [][]float64) *mat.SymDense {

	N := len(X)
	Dim := len(X[0])
	C := mat.NewSymDense(N, nil)
	tl2 := -1 / (2 * ec.L * ec.L)
	sf2 := ec.SigmaF * ec.SigmaF
	for i := 0; i < N; i++ {
		vi := mat.NewVecDense(Dim, X[i])
		for j := i; j < N; j++ {
			if i == j {
				C.SetSym(i, i, sf2+ec.SigmaN*ec.SigmaN)
			} else {
				vj := mat.NewVecDense(Dim, X[j])
				vj.SubVec(vi, vj)
				v := math.Exp(mat.Dot(vj, vj)*tl2) * sf2
				C.SetSym(i, j, v)
			}
		}
	}
	return C
}

func (ec *ExpCov) CalcASymCov(X1 [][]float64, X2 [][]float64) *mat.Dense {

	N1 := len(X1)
	N2 := len(X2)
	Dim := len(X1[0])
	C := mat.NewDense(N1, N2, nil)
	tl2 := -1 / (2 * ec.L * ec.L)
	sf2 := ec.SigmaF * ec.SigmaF
	for i := 0; i < N1; i++ {
		vi := mat.NewVecDense(Dim, X1[i])
		for j := 0; j < N2; j++ {
			vj := mat.NewVecDense(Dim, X2[j])
			vj.SubVec(vj, vi)
			v := math.Exp(mat.Dot(vj, vj)*tl2) * sf2
			C.Set(i, j, v)
		}
	}
	return C
}
