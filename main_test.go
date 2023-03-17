package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.com/chadgh/genetic/genetic"
	"gitlab.com/chadgh/genetic/queens"
)

func Test_queens(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	strategy := genetic.NewGenericStrategy(
		8,    // organism size
		20,   // population size
		0.99, // rate to select the highest fit parents
		16.0, // fitness target
		3000, // generation limit
		0.05, // rate of mutation
		genetic.GenerateAlphabet([]int{0, 1, 2, 3, 4, 5, 6, 7}), // alphabet for the DNA
		queens.Fitness, // fitness function
	)

	winner, _ := strategy.Evolve()
	assert.Equal(t, 16.0, winner.Fitness)
}
