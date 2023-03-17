package projectestimation

import (
	"fmt"

	"gitlab.com/chadgh/genetic/genetic"
)

func Run() {
	strategy := genetic.NewGenericStrategy(
		5,
		10,
		0.5,
		270000.0,
		100,
		0.0,
		[]byte{byte(0), byte(1)},
		func(o genetic.Organism) float64 { return 1.0 },
	)
	winner, generations := strategy.Evolve()
	fmt.Println("Wining solution is will cost: $", cost(winner))
	fmt.Println("The projects to do are: fitness:", winner.Fitness, " generations:", generations)
	for i := range winner.DNA {
		if winner.DNA[i] == 1 {
			fmt.Println("\t", fitness.projects[i].name)
		}
	}
}

func cost(o genetic.Organism) float64 {
	cost := 0.0
	for i := range o.DNA {
		if o.DNA[i] == 1 {
			cost = cost + fitness.projects[i].cost
		}
	}
	return cost
}

var fitness = myfitness{
	projects: []Project{
		{
			name:    "External entrance",
			cost:    20000.0,
			benefit: 5.0,
		},
		{
			name:    "Pattio",
			cost:    10000.0,
			benefit: 7.0,
		},
		{
			name:    "Backyard stairs",
			cost:    5000.0,
			benefit: 7.0,
		},
		{
			name:    "Fence",
			cost:    7000.0,
			benefit: 3.0,
		},
		{
			name:    "Furnace",
			cost:    12000.0,
			benefit: 6.0,
		},
	},
}
