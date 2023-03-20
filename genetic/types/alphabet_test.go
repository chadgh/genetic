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

func TestNewIntAlphabet(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want Alphabet
	}{
		{
			name: "only 0, 1",
			args: args{min: 0, max: 1},
			want: Alphabet([]byte{0, 1}),
		},
		{
			name: "0, 10",
			args: args{min: 0, max: 10},
			want: Alphabet([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
		{
			name: "1, 10",
			args: args{min: 1, max: 10},
			want: Alphabet([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewIntAlphabet(tt.args.min, tt.args.max); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIntAlphabet() = %v, want %v", got, tt.want)
			}
		})
	}
}
