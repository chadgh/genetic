package catdog

import (
	"log"

	"gitlab.com/chadgh/genetic/genetic"
	"gitlab.com/chadgh/genetic/genetic/strategies"
	genetictypes "gitlab.com/chadgh/genetic/genetic/types"
)

func Run() {
	alphabet := genetictypes.NewIntAlphabet(0, 42)
	strategy := strategies.NewBasicStrategy(
		alphabet,
		321.0,
		10000,
		0.05,
		Fitness,
	)
	winner, generations := genetic.RunGenerations(
		100,
		2,
		alphabet,
		strategy,
	)

	log.Println("dna:", winner.DNA)
	log.Println("fitness:", strategy.Fitness(winner))
	log.Println("generations:", generations)

	cats, dogs := int(winner.DNA[0]), int(winner.DNA[1])
	log.Println("startup cost: $", StartUpCost(cats, dogs))
	log.Println("sqft used: ", TotalSqftUsed(cats, dogs), " sqft")
	timeNeeded := TotalTimeNeeded(cats, dogs)
	log.Println("time needed: ", timeNeeded, " min OR ", timeNeeded/60, " hours ", timeNeeded%60, " min")
	log.Println("weekly earnings: $", TotalEarnings(cats, dogs))
}
