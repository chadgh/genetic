package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/chadgh/genetic/genetic"
	"gitlab.com/chadgh/genetic/queens"
)

func Test_queens(t *testing.T) {
	strategy := genetic.NewGenericStrategy(
		8,    // organism size
		100,  // population size
		0.90, // rate to select the highest fit parents
		16.0, // fitness target
		3000, // generation limit
		0.25, // rate of mutation
		genetic.GenerateAlphabet([]int{0, 1, 2, 3, 4, 5, 6, 7}), // alphabet for the DNA
		queens.Fitness, // fitness function
	)

	winner, _ := strategy.Evolve()
	assert.Equal(t, 16.0, winner.Fitness)
}
