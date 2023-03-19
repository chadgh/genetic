package types

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewRandomOrganism(t *testing.T) {
	rand.Seed(1)
	type args struct {
		size     int
		alphabet Alphabet
	}
	tests := []struct {
		name string
		args args
		want Organism
	}{
		{
			name: "simple",
			args: args{
				size:     8,
				alphabet: NewAlphabet([]int{1, 2, 3, 4}),
			},
			want: Organism{
				DNA:     []byte{2, 4, 4, 4, 2, 3, 2, 1},
				Fitness: 0.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomOrganism(tt.args.size, tt.args.alphabet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomOrganism() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEmptyOrganism(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want Organism
	}{
		{
			name: "simple",
			args: args{size: 6},
			want: Organism{
				DNA:     Alphabet(make([]byte, 6)),
				Fitness: 0.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyOrganism(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyOrganism() = %v, want %v", got, tt.want)
			}
		})
	}
}
