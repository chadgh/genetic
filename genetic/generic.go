package genetic

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type GenericStrategy struct {
	organismSize       int
	populationSize     int
	selectionThreshold float64
	fitnessTarget      float64
	generationLimit    int
	mutationRate       float64
	alphabet           []byte
	fitnessFunc        FitnessFunc
}

func NewGenericStrategy(
	organismSize int,
	populationSize int,
	selectionThreshold float64,
	fitnessTarget float64,
	generationLimit int,
	mutationRate float64,
	alphabet []byte,
	fitnessFunc FitnessFunc,
) GenericStrategy {
	return GenericStrategy{
		organismSize:       organismSize,
		populationSize:     populationSize,
		selectionThreshold: selectionThreshold,
		fitnessTarget:      fitnessTarget,
		generationLimit:    generationLimit,
		mutationRate:       mutationRate,
		alphabet:           alphabet,
		fitnessFunc:        fitnessFunc,
	}
}

func (g *GenericStrategy) GenerationLimit() int {
	return g.generationLimit
}

func (g *GenericStrategy) FitnessTarget() float64 {
	return g.fitnessTarget
}

func (g *GenericStrategy) NewRandomOrganism(size int) Organism {
	dna := make([]byte, size)
	for i := range dna {
		dna[i] = byte(g.alphabet[rand.Intn(len(g.alphabet))])
	}
	o := Organism{DNA: dna}
	o.Fitness = g.fitnessFunc(o)
	return o
}

func (g *GenericStrategy) Populate() []Organism {
	population := make([]Organism, g.populationSize)
	for p := range population {
		population[p] = g.NewRandomOrganism(g.organismSize)
	}
	return population
}

func (g *GenericStrategy) Select(population []Organism, number int) []Organism {
	selection := make([]Organism, number)
	for n := range selection {
		val := rand.Float64()
		if val <= g.selectionThreshold {
			selection[n] = g.Highest(population, number)[n]
		} else {
			s := rand.Intn(len(population))
			selection[n] = population[s]
		}
	}
	return selection
}

func (g *GenericStrategy) Reproduce(o1, o2 Organism) Organism {
	if len(o1.DNA) != len(o2.DNA) || len(o1.DNA) != g.organismSize {
		log.Println(fmt.Sprintf("ERROR: Reproduce random child, bad Organisms: o1 len: %v, o2 len: %v, expected: %v", len(o1.DNA), len(o2.DNA), g.organismSize))
		return g.NewRandomOrganism(g.organismSize)
	}

	c := rand.Intn(g.organismSize)
	child := Organism{}
	child.DNA = append(child.DNA, o1.DNA[:c]...)
	child.DNA = append(child.DNA, o2.DNA[c:]...)

	return child
}

func (g *GenericStrategy) Mutate(o Organism) Organism {
	if rand.Float64() <= g.mutationRate {
		mutationIndex := rand.Intn(len(o.DNA))
		currentValueIndex := bytes.IndexByte(g.alphabet, o.DNA[mutationIndex])
		var mutationValues []byte
		if currentValueIndex > 0 && currentValueIndex < len(g.alphabet) {
			mutationValues = append(g.alphabet[0:currentValueIndex], g.alphabet[currentValueIndex+1:]...)
		} else if currentValueIndex == 0 {
			mutationValues = g.alphabet[1:]
		} else {
			mutationValues = g.alphabet[0:len(g.alphabet)]
		}
		// if len(mutationValues) != len(g.alphabet)-1 {
		// 	panic(fmt.Sprintf("Something is wrong with the mutation values! %d vs %d", len(mutationValues), len(g.alphabet)-1))
		// }
		mutationValue := mutationValues[rand.Intn(len(mutationValues))]
		o.DNA[mutationIndex] = mutationValue
	}
	return o
}

func (g *GenericStrategy) CalcFitness(population []Organism) {
	for p := range population {
		population[p].Fitness = g.fitnessFunc(population[p])
	}
}

func (g *GenericStrategy) Evolve() (Organism, int) {
	found := false
	generation := 0
	var bestFitOrganism Organism
	population := g.Populate()
	g.CalcFitness(population)
	for !found {
		generation++
		newPopulation := make([]Organism, len(population))
		for i := 0; i < len(population); i++ {
			parents := g.Select(population, 2)
			child := g.Reproduce(parents[0], parents[1])
			child = g.Mutate(child)
			newPopulation[i] = child
		}
		population = newPopulation
		if len(population) != g.populationSize {
			panic("Population size fail in evolve")
		}
		g.CalcFitness(population)

		fitOrganism := g.Highest(population, 1)[0]
		if fitOrganism.Fitness >= bestFitOrganism.Fitness {
			bestFitOrganism = fitOrganism
		}
		if bestFitOrganism.Fitness >= g.fitnessTarget {
			break
		}
		if generation == g.GenerationLimit() {
			break
		}
	}
	return bestFitOrganism, generation
}

func (g *GenericStrategy) Highest(population []Organism, number int) []Organism {
	highest := make([]Organism, number)
	for h := range highest {
		highestIndex := 0
		highest[h] = population[highestIndex]
		for o := range population {
			if population[o].Fitness > highest[h].Fitness {
				highest[h] = population[o]
				highestIndex = o
			}
		}
		population = append(population[:highestIndex], population[highestIndex+1:]...)
	}
	return highest
}
