package catdog

import (
	"fmt"

	"gitlab.com/chadgh/genetic/genetic"
)

func Run() {
	strategy := genetic.NewGenericStrategy(
		10,
		10,
		0.75,
		30000.0,
		1000,
		0.5,
		[]byte{},
		fitness,
	)
	winner, _ := strategy.Evolve()
	fmt.Println(winner)
}

var fitness = myfitness{}
