package genetic

type Organism struct {
	DNA     []byte
	Fitness float64
}

type FitnessFunc = func(Organism) float64

type Alphabet = []byte
