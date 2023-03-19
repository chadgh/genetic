package types

type Probs = []float64

func NewProbs(size int) Probs {
	return Probs(make([]float64, size))
}
