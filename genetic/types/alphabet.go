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
