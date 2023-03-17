package napsack

import (
	"fmt"

	"gitlab.com/chadgh/genetic/genetic"
)

type Thing struct {
	name   string
	weight float64
	value  float64
}
type myfitness struct {
	things []Thing
	limit  float64
}

func (f myfitness) Fitness(organism genetic.Organism) float64 {
	weight := 0.0
	value := 0.0
	fmt.Println("length", len(organism.DNA))
	if len(organism.DNA) != len(f.things) {
		return 0.0
	}
	for i := 0; i < len(f.things); i++ {
		if organism.DNA[i] == 1 {
			weight += f.things[i].weight
			value += f.things[i].value

			if weight > f.limit {
				return 0.0
			}
		}

	}
	return value
}
