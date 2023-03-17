package main

import (
	"bytes"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/chadgh/genetic/genetic"
	"gitlab.com/chadgh/genetic/queens"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	return buf.String()
}

func Test_queens(t *testing.T) {
	// rand.Seed(time.Now().UTC().UnixNano())
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
	// output := captureOutput(func() { queens.Run() })
	// fitness, err := strconv.Atoi(strings.Trim(strings.Split(strings.Split(output, "\n")[1], ":")[3], " "))
	// assert.Equal(t, nil, err, "no fitness")
	// assert.Equal(t, 16, fitness, "They don't match.")
	// assert.Fail(t, output)
	assert.Equal(t, 16.0, winner.Fitness)
}
