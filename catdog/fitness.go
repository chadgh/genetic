package catdog

import (
	"github.com/chadgh/genetic/genetic/types"
)

type Costs struct {
	initialCatCost float64
	initialDogCost float64
	initialAmount  float64
	catSqft        int
	dogSqft        int
	maxSqft        int
	catTime        int
	dogTime        int
	maxTime        int
	catPay         float64
	dogPay         float64
	maxPay         float64
}

var scenario = Costs{
	30.0,   // initialCatCost float64
	80.0,   // initialDogCost float64
	1280.0, // initialAmount float64
	8,      // catSqft int
	24,     // dogSqft int
	360,    // maxSqft int
	16,     // catTime int
	20,     // dogTime int
	480,    // maxTime int
	8.0,    // catPay float64
	20.0,   // dogPay float64
	1000.0, // maxPay float64
}

func Fitness(organism types.Organism) float64 {
	cats, dogs := int(organism.DNA[0]), int(organism.DNA[1])
	initialCost := StartUpCost(cats, dogs)
	sqftCost := TotalSqftUsed(cats, dogs)
	timeCost := TotalTimeNeeded(cats, dogs)
	totalPay := TotalEarnings(cats, dogs)

	if initialCost <= scenario.initialAmount &&
		sqftCost <= scenario.maxSqft &&
		timeCost <= scenario.maxTime {
		return totalPay
	}
	return 0.0
}

func StartUpCost(cats, dogs int) float64 {
	return (float64(cats) * scenario.initialCatCost) + (float64(dogs) * scenario.initialDogCost)
}

func TotalSqftUsed(cats, dogs int) int {
	return (cats * scenario.catSqft) + (dogs * scenario.dogSqft)
}

func TotalTimeNeeded(cats, dogs int) int {
	return (cats * scenario.catTime) + (dogs * scenario.dogTime)
}

func TotalEarnings(cats, dogs int) float64 {
	return (float64(cats) * scenario.catPay) + (float64(dogs) * scenario.dogPay)
}
