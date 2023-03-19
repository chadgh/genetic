package types

import (
	"reflect"
	"testing"
)

func TestNewAlphabet(t *testing.T) {
	t.Run("standard int", func(t *testing.T) {
		intItems := []int{1, 2, 3, 4}
		want := Alphabet([]byte{1, 2, 3, 4})
		if got := NewAlphabet(intItems); !reflect.DeepEqual(got, want) {
			t.Errorf("NewAlphabet(int) = %v, want %v", got, want)
		}
	})
	t.Run("int64", func(t *testing.T) {
		intItems := []int64{1, 2, 3, 4}
		want := Alphabet([]byte{1, 2, 3, 4})
		if got := NewAlphabet(intItems); !reflect.DeepEqual(got, want) {
			t.Errorf("NewAlphabet(int64) = %v, want %v", got, want)
		}
	})
	t.Run("runes", func(t *testing.T) {
		intItems := []rune{'a', 'b', 'c', 'd'}
		want := Alphabet([]byte{97, 98, 99, 100})
		if got := NewAlphabet(intItems); !reflect.DeepEqual(got, want) {
			t.Errorf("NewAlphabet(runes) = %v, want %v", got, want)
		}
	})
}
