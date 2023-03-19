package types

type Strategy interface {
	Probability(Organism) float64
	Selection(Population, Probs) Organism
	Reproduce(o1, o2 Organism) Organism
	Mutate(o Organism) Organism
	Fitness(o Organism) float64
	MaxFitness() float64
	MaxGenerations() int
}
