package queens

import (
	"log"

	"gitlab.com/chadgh/genetic/genetic"
)

func Run() {
	strategy := genetic.NewGenericStrategy(
		8,      // organism size
		100,    // population size
		0.97,   // rate to select the highest fit parents
		16.0,   // fitness target
		100000, // generation limit
		0.5,    // rate of mutation
		genetic.GenerateAlphabet([]int{0, 1, 2, 3, 4, 5, 6, 7}), // alphabet for the DNA
		Fitness, // fitness function
	)

	// strategy := genetic.NewGenericStrategy(
	// 	5,
	// 	100,
	// 	0.99,
	// 	5.0,
	// 	10000,
	// 	0.05,
	// 	genetic.GenerateAlphabet([]int{0, 1, 2, 3, 4}),
	// 	Fitness,
	// )

	winner, generations := strategy.Evolve()
	log.Println("dna:", winner.DNA)
	log.Println("fitness:", winner.Fitness)
	log.Println("generations:", generations)
}
