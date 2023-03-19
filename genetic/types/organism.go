package types

import "math/rand"

type Organism struct {
	DNA     Alphabet
	Fitness float64
}

func NewRandomOrganism(size int, alphabet Alphabet) Organism {
	organism := Organism{
		DNA:     Alphabet(make([]byte, size)),
		Fitness: 0.0,
	}
	for o := range organism.DNA {
		organism.DNA[o] = alphabet[rand.Intn(len(alphabet))]
	}
	return organism
}

func NewEmptyOrganism(size int) Organism {
	organism := Organism{
		DNA:     Alphabet(make([]byte, size)),
		Fitness: 0.0,
	}
	return organism
}
