package genetic

type Alphabetic interface {
	~int | ~rune
}

func GenerateAlphabet[T Alphabetic](stuff []T) Alphabet {
	alphabet := make([]byte, len(stuff))
	for i := range stuff {
		alphabet[i] = byte(stuff[i])
	}
	return alphabet
}
