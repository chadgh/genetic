package types

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestNewRandomPopulation(t *testing.T) {
	rand.Seed(1)
	type args struct {
		populationSize int
		organismSize   int
		alphabet       Alphabet
	}
	tests := []struct {
		name string
		args args
		want Population
	}{
		{
			name: "basic",
			args: args{
				populationSize: 3,
				organismSize:   4,
				alphabet:       []byte{1, 2, 3, 4},
			},
			want: Population{
				{DNA: []byte{2, 4, 4, 4}, Fitness: 0.0},
				{DNA: []byte{2, 3, 2, 1}, Fitness: 0.0},
				{DNA: []byte{1, 1, 3, 4}, Fitness: 0.0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRandomPopulation(tt.args.populationSize, tt.args.organismSize, tt.args.alphabet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRandomPopulation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopulation_Sort(t *testing.T) {
	tests := []struct {
		name string
		p    Population
		want Population
	}{
		{
			name: "basic",
			p: Population{
				{Fitness: 4},
				{Fitness: 3},
				{Fitness: 10},
				{Fitness: 4},
			},
			want: Population{
				{Fitness: 3},
				{Fitness: 4},
				{Fitness: 4},
				{Fitness: 10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Sort()
			if !reflect.DeepEqual(tt.p, tt.want) {
				t.Errorf("Sort() = %v, want %v", tt.p, tt.want)
			}
		})
	}
}
