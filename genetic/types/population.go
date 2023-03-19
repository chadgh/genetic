package types

import "sort"

type Population []Organism

// NewPopulation creates a new population of the specified size
// with empty organisms
func NewPopulation(size int) Population {
	return Population(make([]Organism, size))
}

// NewRandomPopulation creates a new population of the specififed size
// with random organisms populated
func NewRandomPopulation(populationSize, organismSize int, alphabet Alphabet) Population {
	population := NewPopulation(populationSize)
	for p := range population {
		population[p] = NewRandomOrganism(organismSize, alphabet)
	}
	return population
}

// Population functions to make it sortable
func (p Population) Len() int {
	return len(p)
}

func (p Population) Less(i, j int) bool {
	iValue := p[i].Fitness
	jValue := p[j].Fitness
	return iValue < jValue
}

func (p Population) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Population) Sort() {
	sort.Sort(p)
}
