package napsack

import (
	"fmt"

	"gitlab.com/chadgh/genetic/genetic"
)

func Run() {
	strategy := genetic.NewGenericStrategy(
		5,
		10,
		0.5,
		750.0,
		100,
		0.5,
		[]byte{byte(0), byte(1)},
		func(o genetic.Organism) float64 { return 1.0 },
	)

	winner, generations := strategy.Evolve()
	fmt.Println(winner, generations)
}

var fitness = myfitness{
	things: []Thing{
		{
			name:   "Laptop",
			value:  500.0,
			weight: 2200.0,
		},
		{
			name:   "Headphones",
			value:  150.0,
			weight: 160.0,
		},
		{
			name:   "Coffee Mug",
			value:  60.0,
			weight: 350.0,
		},
		{
			name:   "Notepad",
			value:  40.0,
			weight: 333.0,
		},
		{
			name:   "Water Bottle",
			value:  30.0,
			weight: 192.0,
		},
	},
	limit: 3000.0,
}
