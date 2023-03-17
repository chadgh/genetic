package genetic

import "sort"

type Organism struct {
	DNA     []byte
	Fitness float64
}

type FitnessFunc = func(Organism) float64

type Alphabet = []byte

type Population []Organism

func NewPopulation(size int) Population {
	return Population(make([]Organism, size))
}

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
