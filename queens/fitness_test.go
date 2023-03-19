package queens

import (
	"testing"

	"gitlab.com/chadgh/genetic/genetic/types"
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
			name: "same rank",
			args: args{
				types.Organism{
					DNA: []byte{
						byte(0),
						byte(0),
						byte(0),
						byte(0),
						byte(0),
						byte(0),
						byte(0),
						byte(0),
					},
				},
			},
			want: 1.0,
		},
		{
			name: "different rank, still bad",
			args: args{
				types.Organism{
					DNA: []byte{
						byte(0),
						byte(1),
						byte(2),
						byte(3),
						byte(4),
						byte(5),
						byte(6),
						byte(7),
					},
				},
			},
			want: 8.0,
		},
		{
			name: "actual solution #1",
			args: args{
				types.Organism{
					DNA: []byte{
						byte(1),
						byte(3),
						byte(5),
						byte(7),
						byte(2),
						byte(0),
						byte(6),
						byte(4),
					},
				},
			},
			want: 16.0,
		},
		{
			name: "actual solution #2",
			args: args{
				types.Organism{
					DNA: []byte{
						byte(0),
						byte(6),
						byte(3),
						byte(5),
						byte(7),
						byte(1),
						byte(4),
						byte(2),
					},
				},
			},
			want: 16.0,
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
