package projectestimation

import (
	"gitlab.com/chadgh/genetic/genetic"
)

// too many projects not enough money
type Project struct {
	name    string
	cost    float64
	benefit float64
}
type myfitness struct {
	projects []Project
}

func (f myfitness) Fitness(organism genetic.Organism) float64 {
	cost := 0.0
	benefit := 0.0
	for i := 0; i < len(organism.DNA); i++ {
		if organism.DNA[i] == 1 {
			cost = cost + f.projects[i].cost
			benefit = benefit + f.projects[i].benefit
		}
	}
	benefit = benefit / float64(len(organism.DNA))
	if cost > 30000.0 {
		cost = 0.0
	}
	return cost * benefit
}
