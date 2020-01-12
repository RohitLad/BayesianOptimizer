package kernel

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type ExpCov struct {
	SigmaF     float64
	L          float64
	SigmaN     float64
	factorsset bool
	tl2        float64
	sf2        float64
	sn2        float64
}

func NewExpCov(SigmaF, L, SigmaN float64) *ExpCov {
	cov := ExpCov{SigmaF: SigmaF, L: L, SigmaN: SigmaN}
	cov.SetFactors()
	return &cov
}

func (ec *ExpCov) CalcSymCov(X [][]float64) *mat.SymDense {

	N := len(X)
	Dim := len(X[0])
	C := mat.NewSymDense(N, nil)

	for i := 0; i < N; i++ {
		vi := mat.NewVecDense(Dim, X[i])
		for j := i; j < N; j++ {
			if i == j {
				C.SetSym(i, i, ec.sf2+ec.sn2)
			} else {
				vj := mat.NewVecDense(Dim, X[j])
				C.SetSym(i, j, ec.calcExponent(vi, vj))
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

	for i := 0; i < N1; i++ {
		vi := mat.NewVecDense(Dim, X1[i])
		for j := 0; j < N2; j++ {
			vj := mat.NewVecDense(Dim, X2[j])
			C.Set(i, j, ec.calcExponent(vi, vj))
		}
	}
	return C
}

func (ec *ExpCov) SetFactors() {
	if !ec.factorsset {
		ec.sf2 = ec.SigmaF * ec.SigmaF
		ec.tl2 = -0.5 / (ec.L * ec.L)
		ec.sn2 = ec.SigmaN * ec.SigmaN
	}
}

func (ec *ExpCov) calcExponent(v1, v2 *mat.VecDense) float64 {
	v2.SubVec(v1, v2)
	return math.Exp(mat.Dot(v2, v2)*ec.tl2) * ec.sf2
}
