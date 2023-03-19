package eightqueens

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var numQueens = 8
var maxFitness = float64((numQueens * (numQueens - 1)) / 2.0)
var maxFitness2 = 16.0
var mutationProbability = 0.03

type Organism []int

// init sets up the package by:
// * setting a random seed
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// RandomOrganism creates and returns a random array of integers
func RandomOrganism(size int) Organism {
	organism := Organism(make([]int, size))
	for o := range organism {
		organism[o] = rand.Intn(numQueens) + 1
	}
	return organism
}

// Fitness measures the fitness value of an organism
// Returns the fitness value or the "fit" value of an organism
func Fitness(organism Organism) int {
	// find horizontal collisions and left and right diagonal collisions
	files := map[int]bool{}
	leftDiagonal := make([]int, 2*numQueens)
	rightDiagonal := make([]int, 2*numQueens)
	for i := 0; i < len(organism); i++ {
		files[int(organism[i])] = true
		leftIndex := i + organism[i] - 1
		rightIndex := len(organism) - i + organism[i] - 2
		leftDiagonal[leftIndex]++
		rightDiagonal[rightIndex]++
	}
	horizontalCollisions := numQueens - len(files)

	// find total diagonal collisions
	diagonalCollisions := 0
	for i := 0; i < 2*numQueens-1; i++ {
		counter := 0
		if leftDiagonal[i] > 1 {
			counter = counter + leftDiagonal[i] - 1
		}
		if rightDiagonal[i] > 1 {
			counter = counter + rightDiagonal[i] - 1
		}
		diagonalCollisions = diagonalCollisions + (counter / (numQueens - int(math.Abs(float64(i-numQueens+1)))))

	}

	// Fitness score is the max fitness minus any collisions
	return int(maxFitness) - (horizontalCollisions + diagonalCollisions)
}

func Fitness2(organism Organism) int {
	files := map[int]bool{}
	ranks := organism
	for q := range ranks {
		files[int(ranks[q])] = true
	}
	score := len(files) // max is 8 at this point

	if len(files) != len(ranks) {
		return score
	}

	// diagonal check
	for q1 := range ranks {
		hit := false
		for q2 := range ranks {
			if q1 == q2 {
				continue
			}
			diff := int(math.Abs(float64(q1 - q2)))
			v2down := int(ranks[q1]) - diff
			v2up := int(ranks[q1]) + diff
			if v2down >= 0 && int(ranks[q2]) == v2down {
				hit = true
				break
			}
			if v2up <= 7 && int(ranks[q2]) == v2up {
				hit = true
				break
			}
		}
		if !hit {
			score++
		}
	}
	return score // max score should be 16
}

// Probability returns the probability of an organism being picked
// This is kind of how "good" the organism is
func Probability(organism Organism) float64 {
	return float64(Fitness(organism)) / maxFitness
}

func Probability2(organism Organism) float64 {
	return float64(Fitness2(organism)) / maxFitness2
}

type Pair struct {
	organism    Organism
	probability float64
}

// RandomSelection returns a random selection from the population based on the probabilities
// of seeing the organism
func RandomSelection(population []Organism, probabilities []float64) Organism {
	pairs := make([]Pair, len(population))
	total := 0.0
	for i := 0; i < len(pairs); i++ {
		pairs[i] = Pair{organism: population[i], probability: probabilities[i]}
		total += pairs[i].probability
	}
	selectionValue := rand.Float64()
	upto := 0.0
	for i := 0; i < len(pairs); i++ {
		if upto+pairs[i].probability >= selectionValue {
			return pairs[i].organism
		}
		upto += pairs[i].probability
	}
	return population[0]
}

// Reproduce does a crossover between two organisms
// Returns the new organism
func Reproduce(x, y Organism) Organism {
	crossover := rand.Intn(numQueens)
	return append(x[0:crossover], y[crossover:numQueens]...)
}

// Mutate changes a random chromosome in an organism
func Mutate(x Organism) Organism {
	mutationIndex := rand.Intn(numQueens)
	mutationValue := rand.Intn(numQueens) + 1
	x[mutationIndex] = mutationValue
	return x
}

func Evolv(population []Organism) []Organism {
	newPopulation := make([]Organism, len(population))
	probabilities := make([]float64, len(population))
	for i := range probabilities {
		probabilities[i] = Probability2(population[i])
	}
	for i := range population {
		x := RandomSelection(population, probabilities)
		y := RandomSelection(population, probabilities)
		child := Reproduce(x, y)
		if rand.Float64() < mutationProbability {
			child = Mutate(child)
		}
		newPopulation[i] = child
		if Fitness2(child) == int(maxFitness2) {
			break
		}
	}
	return newPopulation
}

func Run() {
	population := make([]Organism, 100)
	for i := range population {
		population[i] = RandomOrganism(numQueens)
	}
	generation := 1

	var found *Organism = nil

	for {
		for i := range population {
			if Fitness2(population[i]) == int(maxFitness2) {
				found = &population[i]
				break
			}
		}
		if found != nil {
			break
		}
		population = Evolv(population)
		generation++
	}

	fmt.Printf("Solved in generation %v!\n", generation)
	fmt.Printf("Solution:\n\n%v\n\n", *found)
}
