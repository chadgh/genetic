package catdog

import (
	"testing"

	"github.com/chadgh/genetic/genetic/types"
)

func TestFitness(t *testing.T) {
	type args struct {
		organism types.Organism
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "solution 1",
			args: args{organism: types.Organism{DNA: []byte{10, 10}}},
			want: 210.0,
		},
		{
			name: "non-solution 1 - too many dogs",
			args: args{organism: types.Organism{DNA: []byte{0, 17}}},
			want: 0.0,
		},
		{
			name: "non-solution 2 - too many cats",
			args: args{organism: types.Organism{DNA: []byte{43, 0}}},
			want: 0.0,
		},
		{
			name: "non-solution 3 - too expensive",
			args: args{organism: types.Organism{DNA: []byte{10, 15}}},
			want: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fitness(tt.args.organism); got != tt.want {
				t.Errorf("Fitness() = %v, want %v", got, tt.want)
			}
		})
	}
}
