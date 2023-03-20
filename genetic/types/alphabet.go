package types

type Alphabet = []byte

type Alphabetic interface {
	~int | ~int64 | ~rune
}

// NewAlphabet returns a new Alphabet containing the items provided
func NewAlphabet[A Alphabetic](items []A) Alphabet {
	alphabet := make([]byte, len(items))
	for i := range items {
		alphabet[i] = byte(items[i])
	}
	return alphabet
}

func NewIntAlphabet(min, max int) Alphabet {
	size := max - min + 1
	alphabet := make([]byte, size)
	index := 0
	for value := min; value <= max; value++ {
		alphabet[index] = byte(value)
		index++
	}
	return Alphabet(alphabet)
}
