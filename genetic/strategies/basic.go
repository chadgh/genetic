package strategies

import (
	"math/rand"

	"github.com/chadgh/genetic/genetic/types"
)

type BasicStrategy struct {
	alphabet            types.Alphabet
	maxFitness          float64
	maxGenerations      int
	mutationProbability float64
	fitnessFunc         types.FitnessFunc
}

func NewBasicStrategy(
	alphabet types.Alphabet,
	maxFitness float64,
	maxGenerations int,
	mutationProbability float64,
	fitnessFunc types.FitnessFunc,
) BasicStrategy {
	return BasicStrategy{
		alphabet:            alphabet,
		maxFitness:          maxFitness,
		maxGenerations:      maxGenerations,
		mutationProbability: mutationProbability,
		fitnessFunc:         fitnessFunc,
	}
}

func (s BasicStrategy) Probability(organism types.Organism) float64 {
	return s.Fitness(organism) / s.maxFitness
}

// Selection does a random selection based on the probabilities from the population
func (s BasicStrategy) Selection(population types.Population, probabilities types.Probs) types.Organism {
	type Pair struct {
		o types.Organism
		p float64
	}
	popSize := len(population)
	zip := make([]Pair, popSize)
	total := 0.0
	for i := range zip {
		zip[i] = Pair{o: population[i], p: probabilities[i]}
		total += zip[i].p
	}
	selectionValue := rand.Float64()
	upto := 0.0
	for i := range zip {
		if upto+zip[i].p >= selectionValue {
			return zip[i].o
		}
		upto += zip[i].p
	}
	return population[0]
}

func (s BasicStrategy) Reproduce(o1, o2 types.Organism) types.Organism {
	orgSize := len(o1.DNA)
	crossover := rand.Intn(orgSize)
	child := types.NewEmptyOrganism(orgSize)
	child.DNA = o1.DNA[0:crossover]
	child.DNA = append(child.DNA, o2.DNA[crossover:orgSize]...)
	return child
}

func (s BasicStrategy) Mutate(o types.Organism) types.Organism {
	if rand.Float64() < s.mutationProbability {
		mutationIndex := rand.Intn(len(o.DNA))
		mutationValue := s.alphabet[rand.Intn(len(s.alphabet))]
		o.DNA[mutationIndex] = mutationValue
	}
	return o
}

func (s BasicStrategy) Fitness(o types.Organism) float64 {
	return s.fitnessFunc(o)
}

func (s BasicStrategy) MaxFitness() float64 {
	return s.maxFitness
}

func (s BasicStrategy) MaxGenerations() int {
	return s.maxGenerations
}
