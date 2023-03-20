package strategies

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/chadgh/genetic/genetic/types"
)

func TestBasicStrategy_Mutate(t *testing.T) {
	rand.Seed(1)
	type fields struct {
		alphabet            types.Alphabet
		maxFitness          float64
		maxGenerations      int
		mutationProbability float64
	}
	type args struct {
		o types.Organism
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   types.Organism
	}{
		{
			name: "mutation 1",
			fields: fields{
				alphabet:            types.NewAlphabet([]int{1, 2, 3, 4, 5}),
				mutationProbability: 1.0,
			},
			args: args{
				o: types.Organism{DNA: []byte{1, 1, 1}},
			},
			want: types.Organism{DNA: []byte{3, 1, 1}},
		},
		{
			name: "no mutation",
			fields: fields{
				alphabet:            types.NewAlphabet([]int{1, 2, 3, 4, 5}),
				mutationProbability: 0.0,
			},
			args: args{
				o: types.Organism{DNA: []byte{1, 1, 1}},
			},
			want: types.Organism{DNA: []byte{1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BasicStrategy{
				alphabet:            tt.fields.alphabet,
				maxFitness:          tt.fields.maxFitness,
				maxGenerations:      tt.fields.maxGenerations,
				mutationProbability: tt.fields.mutationProbability,
			}
			if got := s.Mutate(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicStrategy.Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicStrategy_Reproduce(t *testing.T) {
	rand.Seed(1)
	type args struct {
		o1 types.Organism
		o2 types.Organism
	}
	tests := []struct {
		name string
		args args
		want types.Organism
	}{
		{
			name: "reproduce 1",
			args: args{
				o1: types.Organism{DNA: []byte{1, 1, 1, 1, 1}},
				o2: types.Organism{DNA: []byte{2, 2, 2, 2, 2}},
			},
			want: types.Organism{DNA: []byte{1, 2, 2, 2, 2}},
		},
		{
			name: "reproduce 2",
			args: args{
				o1: types.Organism{DNA: []byte{1, 1, 1, 1, 1, 1}},
				o2: types.Organism{DNA: []byte{2, 2, 2, 2, 2, 2}},
			},
			want: types.Organism{DNA: []byte{1, 1, 1, 2, 2, 2}},
		},
		{
			name: "reproduce 3",
			args: args{
				o1: types.Organism{DNA: []byte{1, 1, 1}},
				o2: types.Organism{DNA: []byte{2, 2, 2}},
			},
			want: types.Organism{DNA: []byte{1, 1, 2}},
		},
		{
			name: "reproduce 4",
			args: args{
				o1: types.Organism{DNA: []byte{1, 1, 1, 1, 1}},
				o2: types.Organism{DNA: []byte{2, 2, 2, 2, 2}},
			},
			want: types.Organism{DNA: []byte{1, 1, 1, 1, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BasicStrategy{}
			if got := s.Reproduce(tt.args.o1, tt.args.o2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicStrategy.Reproduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBasicStrategy_Selection(t *testing.T) {
	rand.Seed(1)
	type args struct {
		population    types.Population
		probabilities types.Probs
	}
	tests := []struct {
		name string
		args args
		want types.Organism
	}{
		{
			name: "basic",
			args: args{
				population: types.Population{
					{DNA: []byte{1, 1, 1}},
					{DNA: []byte{2, 2, 2}},
					{DNA: []byte{3, 3, 3}},
				},
				probabilities: []float64{0.1, 0.5, 0.2},
			},
			want: types.Organism{DNA: []byte{3, 3, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := BasicStrategy{}
			if got := s.Selection(tt.args.population, tt.args.probabilities); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BasicStrategy.Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}
