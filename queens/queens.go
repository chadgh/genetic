package queens

import (
	"log"

	"gitlab.com/chadgh/genetic/genetic"
)

func Run() {
	// rand.Seed(time.Now().UTC().UnixNano())
	strategy := genetic.NewGenericStrategy(
		8,    // organism size
		40,   // population size
		0.99, // rate to select the highest fit parents
		16,   // fitness target
		5000, // generation limit
		0.60, // rate of mutation
		genetic.GenerateAlphabet([]int{0, 1, 2, 3, 4, 5, 6, 7}), // alphabet for the DNA
		Fitness, // fitness function
	)

	winner, generations := strategy.Evolve()
	log.Println("dna:", winner.DNA)
	log.Println("fitness:", winner.Fitness)
	log.Println("generations:", generations)
}
