package queens

import (
	"log"

	"gitlab.com/chadgh/genetic/genetic"
	"gitlab.com/chadgh/genetic/genetic/strategies"
	genetictypes "gitlab.com/chadgh/genetic/genetic/types"
)

func Run() {
	alphabet := genetictypes.NewAlphabet([]int{1, 2, 3, 4, 5, 6, 7, 8})
	strategy := strategies.NewBasicStrategy(
		alphabet,
		16.0,
		0,
		0.05,
		Fitness,
	)
	winner, generations := genetic.RunGenerations(
		100,
		8,
		alphabet,
		strategy,
	)

	log.Println("dna:", winner.DNA)
	log.Println("fitness:", winner.Fitness)
	log.Println("generations:", generations)
}
