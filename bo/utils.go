package bo

import "math"

type Normal struct {
	Mu    float64
	Sigma float64
}

func (N Normal) PDF(x float64) float64 {
	return math.Exp(-math.Pow((x-N.Mu)/N.Sigma, 2)/2) / (N.Sigma * math.Sqrt(2*math.Pi))
}

func (N Normal) CDF(x float64) float64 {
	return 0.5 * math.Erfc(-(x-N.Mu)/(N.Sigma*math.Sqrt2))
}
