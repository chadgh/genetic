package queens

import (
	"math"

	"github.com/chadgh/genetic/genetic/types"
)

func Fitness(organism types.Organism) float64 {
	files := map[int]bool{}
	ranks := organism.DNA
	for q := range ranks {
		files[int(ranks[q])] = true
	}
	score := float64(len(files)) // max is 8 at this point

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
