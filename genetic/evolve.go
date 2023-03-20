package genetic

import (
	"github.com/chadgh/genetic/genetic/types"
)

func RunGenerations(
	populationSize int,
	organismSize int,
	alphabet types.Alphabet,
	strategy types.Strategy,
) (types.Organism, int) {
	population := types.NewRandomPopulation(
		populationSize, organismSize, alphabet,
	)
	generation := 1

	var found *types.Organism = nil
	maxGenerations := strategy.MaxGenerations()
	maxFitness := strategy.MaxFitness()

	for {
		for i := range population {
			if strategy.Fitness(population[i]) >= maxFitness {
				found = &population[i]
				break
			}
		}
		if found != nil {
			break
		}
		population = EvolveBasic(population, strategy)
		generation++
		if maxGenerations != 0 && generation == maxGenerations {
			break
		}
	}

	if found == nil {
		population.Sort()
		found = &population[0]
	}
	return *found, generation
}

func EvolveBasic(population types.Population, strategy types.Strategy) types.Population {
	popSize := len(population)
	newPopulation := types.NewPopulation(popSize)
	probabilities := types.NewProbs(popSize)
	for i := range probabilities {
		probabilities[i] = strategy.Probability(population[i])
	}
	for i := range population {
		x := strategy.Selection(population, probabilities)
		y := strategy.Selection(population, probabilities)
		child := strategy.Reproduce(x, y)
		child = strategy.Mutate(child)
		newPopulation[i] = child
		child.Fitness = strategy.Fitness(child)
		if child.Fitness >= strategy.MaxFitness() {
			break
		}
	}
	return newPopulation
}
