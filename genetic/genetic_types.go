package genetic

type Strategy interface {
	GenerationLimit() int
	FitnessTarget() float64
	Fitness(o Organism) float64
	Populate() Population
	Select(population Population) Organism
	Reproduce(o1, o2 Organism) Organism
	Mutate(o Organism) Organism
	Evolve() (Organism, int)
}
