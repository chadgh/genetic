package genetic

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestGenericStrategy_Mutate(t *testing.T) {
	type fields struct {
		selectionThreshold float64
		fitnessTarget      float64
		generationLimit    int
		mutationRate       float64
		alphabet           []byte
	}
	type args struct {
		o Organism
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Organism
	}{
		{
			name: "basic 0 and 1",
			fields: fields{
				selectionThreshold: 0.0,
				fitnessTarget:      0.0,
				generationLimit:    0.0,
				mutationRate:       1.0, // Always mutate
				alphabet:           []byte{byte(0), byte(1)},
			},
			args: args{
				o: Organism{
					DNA: []byte{byte(1)},
				},
			},
			want: Organism{
				DNA: []byte{byte(0)},
			},
		},
		{
			name: "basic 1 and 0",
			fields: fields{
				selectionThreshold: 0.0,
				fitnessTarget:      0.0,
				generationLimit:    0.0,
				mutationRate:       1.0, // Always mutate
				alphabet:           []byte{byte(0), byte(1)},
			},
			args: args{
				o: Organism{
					DNA: []byte{byte(0)},
				},
			},
			want: Organism{
				DNA: []byte{byte(1)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				selectionThreshold: tt.fields.selectionThreshold,
				fitnessTarget:      tt.fields.fitnessTarget,
				generationLimit:    tt.fields.generationLimit,
				mutationRate:       tt.fields.mutationRate,
				alphabet:           tt.fields.alphabet,
			}
			if got := g.Mutate(tt.args.o); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenericStrategy.Mutate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenericStrategy_Highest(t *testing.T) {
	type args struct {
		population []Organism
		number     int
	}
	tests := []struct {
		name string
		args args
		want []Organism
	}{
		{
			name: "Simple 1 Highest",
			args: args{
				population: []Organism{
					{
						Fitness: 1,
					},
					{
						Fitness: 2,
					},
					{
						Fitness: 2,
					},
					{
						Fitness: 4,
					},
				},
				number: 1,
			},
			want: []Organism{
				{
					Fitness: 4,
				},
			},
		},
		{
			name: "1 Highest first",
			args: args{
				population: []Organism{
					{
						Fitness: 5,
					},
					{
						Fitness: 2,
					},
					{
						Fitness: 2,
					},
					{
						Fitness: 4,
					},
				},
				number: 1,
			},
			want: []Organism{
				{
					Fitness: 5,
				},
			},
		},
		{
			name: "2 Highest bookends",
			args: args{
				population: []Organism{
					{
						Fitness: 5,
					},
					{
						Fitness: 2,
					},
					{
						Fitness: 2,
					},
					{
						Fitness: 4,
					},
				},
				number: 2,
			},
			want: []Organism{
				{
					Fitness: 5,
				},
				{
					Fitness: 4,
				},
			},
		},
		{
			name: "Just 1 Highest",
			args: args{
				population: []Organism{
					{
						Fitness: 10,
					},
					{
						Fitness: 100,
					},
					{
						Fitness: 1000,
					},
					{
						Fitness: 1,
					},
				},
				number: 1,
			},
			want: []Organism{
				{
					Fitness: 1000,
				},
			},
		},
		{
			name: "2 Highest",
			args: args{
				population: []Organism{
					{
						Fitness: 10,
					},
					{
						Fitness: 100,
					},
					{
						Fitness: 1000,
					},
					{
						Fitness: 1,
					},
				},
				number: 2,
			},
			want: []Organism{
				{
					Fitness: 1000,
				},
				{
					Fitness: 100,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{}
			got := g.Highest(tt.args.population, tt.args.number)
			if len(got) != len(tt.want) {
				t.Errorf("GenericStrategy.Highest() length not right = %v, want %v", len(got), len(tt.want))
			}
			for i := range got {
				if got[i].Fitness != tt.want[i].Fitness {
					t.Errorf("GenericStrategy.Highest() = %v, want %v", got[i].Fitness, tt.want[i].Fitness)
				}
			}
		})
	}
}

func TestGenericStrategy_CalcFitness(t *testing.T) {
	type fields struct {
		fitnessFunc FitnessFunc
	}
	type args struct {
		population []Organism
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		pre    []float64
		post   []float64
	}{
		{
			name: "Constant fitness",
			fields: fields{
				fitnessFunc: func(o Organism) float64 {
					return 1.0
				},
			},
			args: args{
				population: []Organism{
					{},
					{},
					{},
				},
			},
			pre:  []float64{0.0, 0.0, 0.0},
			post: []float64{1.0, 1.0, 1.0},
		},
		{
			name: "Dynamic fitness",
			fields: fields{
				fitnessFunc: func(o Organism) float64 {
					return o.Fitness + 1
				},
			},
			args: args{
				population: []Organism{
					{Fitness: 1.0},
					{Fitness: 2.0},
					{Fitness: 10.0},
				},
			},
			pre:  []float64{1.0, 2.0, 10.0},
			post: []float64{2.0, 3.0, 11.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				fitnessFunc: tt.fields.fitnessFunc,
			}
			for p := range tt.args.population {
				if tt.args.population[p].Fitness != tt.pre[p] {
					t.Errorf("Precondition not correct.")
				}
			}
			g.CalcFitness(tt.args.population)
			for p := range tt.args.population {
				if tt.args.population[p].Fitness != tt.post[p] {
					t.Errorf("Postcondition not correct.")
				}
			}
		})
	}
}

func TestGenericStrategy_NewRandomOrganism(t *testing.T) {
	rand.Seed(1)
	type fields struct {
		alphabet    []byte
		fitnessFunc FitnessFunc
	}
	type args struct {
		size int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Organism
	}{
		{
			name: "size of 1",
			fields: fields{
				alphabet:    []byte{byte(0), byte(1)},
				fitnessFunc: func(o Organism) float64 { return 1.0 },
			},
			args: args{
				size: 1,
			},
			want: Organism{DNA: []byte{byte(1)}, Fitness: 1.0},
		},
		{
			name: "size of 10",
			fields: fields{
				alphabet:    []byte{byte('a'), byte('b')},
				fitnessFunc: func(o Organism) float64 { return 10.0 },
			},
			args: args{
				size: 10,
			},
			want: Organism{
				DNA: []byte{
					byte('b'),
					byte('b'),
					byte('b'),
					byte('b'),
					byte('a'),
					byte('b'),
					byte('a'),
					byte('a'),
					byte('a'),
					byte('a'),
				},
				Fitness: 10.0,
			},
		},
		{
			name: "size of 3",
			fields: fields{
				alphabet: []byte{byte(1), byte(2), byte(3), byte(4), byte(5), byte(6), byte(7)},
				fitnessFunc: func(o Organism) float64 {
					sum := 0.0
					for i := range o.DNA {
						sum += float64(o.DNA[i])
					}
					return sum
				},
			},
			args: args{
				size: 3,
			},
			want: Organism{
				DNA: []byte{
					byte(7),
					byte(5),
					byte(4),
				},
				Fitness: 16.0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				alphabet:    tt.fields.alphabet,
				fitnessFunc: tt.fields.fitnessFunc,
			}
			got := g.NewRandomOrganism(tt.args.size)
			if len(got.DNA) != tt.args.size {
				t.Errorf("GenericStrategy.NewRandomOrganism() size = %v, want %v", len(got.DNA), tt.want)
			}
			if got.Fitness != tt.want.Fitness {
				t.Errorf("GenericStrategy.NewRandomOrganism() Fitness = %v, want %v", got.Fitness, tt.want.Fitness)
			}
			for d := range got.DNA {
				if got.DNA[d] != tt.want.DNA[d] {
					t.Errorf("GenericStrategy.NewRandomOrganism() [%v] = %v, want %v", d, got.DNA[d], tt.want.DNA[d])
				}
			}
		})
	}
}

func TestGenericStrategy_Populate(t *testing.T) {
	rand.Seed(1)
	type fields struct {
		organismSize   int
		populationSize int
		alphabet       []byte
		fitnessFunc    FitnessFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   []Organism
	}{
		{
			name: "basic",
			fields: fields{
				organismSize:   3,
				populationSize: 2,
				alphabet:       []byte{byte(0), byte(1)},
				fitnessFunc: func(o Organism) float64 {
					sum := 0.0
					for i := range o.DNA {
						sum += float64(o.DNA[i])
					}
					return sum
				},
			},
			want: []Organism{
				{
					DNA:     []byte{byte(1), byte(1), byte(1)},
					Fitness: 3.0,
				},
				{
					DNA:     []byte{byte(1), byte(1), byte(0)},
					Fitness: 2.0,
				},
			},
		},
		{
			name: "pop size 1",
			fields: fields{
				organismSize:   2,
				populationSize: 1,
				alphabet:       []byte{byte(0), byte(1)},
				fitnessFunc: func(o Organism) float64 {
					return 1.0
				},
			},
			want: []Organism{
				{
					DNA:     []byte{byte(1), byte(0)},
					Fitness: 1.0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				organismSize:   tt.fields.organismSize,
				populationSize: tt.fields.populationSize,
				alphabet:       tt.fields.alphabet,
				fitnessFunc:    tt.fields.fitnessFunc,
			}
			population := g.Populate()
			if len(population) != tt.fields.populationSize {
				t.Errorf("GenericStrategy.Populate() size = %v, want %v", len(population), len(tt.want))
			}
			for p := range population {
				if population[p].Fitness != tt.want[p].Fitness {
					t.Errorf("GenericStrategy.Populate() Fitness #%v = %v, want %v", p, population[p], tt.want[p])
				}
				for d := range population[p].DNA {
					if population[p].DNA[d] != tt.want[p].DNA[d] {
						t.Errorf("GenericStrategy.Populate() #%v DNA #%v = %v, want %v", p, d, population[p], tt.want[p])
					}
				}
			}
		})
	}
}

func TestGenericStrategy_Select(t *testing.T) {
	rand.Seed(1)
	type fields struct {
		selectionThreshold float64
	}
	type args struct {
		population []Organism
		number     int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Organism
	}{
		{
			name: "select highest 2",
			fields: fields{
				selectionThreshold: 1.0,
			},
			args: args{
				population: []Organism{
					{
						DNA:     []byte{byte(0)},
						Fitness: 3.0,
					},
					{
						DNA:     []byte{byte(1)},
						Fitness: 5.0,
					},
					{
						DNA:     []byte{byte(2)},
						Fitness: 4.0,
					},
				},
				number: 2,
			},
			want: []Organism{
				{
					DNA:     []byte{byte(1)},
					Fitness: 5.0,
				},
				{
					DNA:     []byte{byte(2)},
					Fitness: 4.0,
				},
			},
		},
		{
			name: "select random 2",
			fields: fields{
				selectionThreshold: 0.0,
			},
			args: args{
				population: []Organism{
					{
						DNA:     []byte{byte(0)},
						Fitness: 3.0,
					},
					{
						DNA:     []byte{byte(1)},
						Fitness: 5.0,
					},
					{
						DNA:     []byte{byte(2)},
						Fitness: 4.0,
					},
				},
				number: 2,
			},
			want: []Organism{
				{
					DNA:     []byte{byte(2)},
					Fitness: 4.0,
				},
				{
					DNA:     []byte{byte(0)},
					Fitness: 3.0,
				},
			},
		},
		{
			name: "select 50/50 2",
			fields: fields{
				selectionThreshold: 0.07,
			},
			args: args{
				population: []Organism{
					{
						DNA:     []byte{byte(0)},
						Fitness: 3.0,
					},
					{
						DNA:     []byte{byte(1)},
						Fitness: 5.0,
					},
					{
						DNA:     []byte{byte(2)},
						Fitness: 4.0,
					},
				},
				number: 2,
			},
			want: []Organism{
				{
					DNA:     []byte{byte(1)},
					Fitness: 5.0,
				},
				{
					DNA:     []byte{byte(2)},
					Fitness: 4.0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				selectionThreshold: tt.fields.selectionThreshold,
			}
			got := g.Select(tt.args.population, tt.args.number)
			if len(got) != tt.args.number {
				t.Errorf("GenericStrategy.Select() size = %v, want %v", len(got), tt.args.number)
			}
			for p := range got {
				if got[p].Fitness != tt.want[p].Fitness {
					t.Errorf("GenericStrategy.Select() Fitness = %v, want %v", got[p].Fitness, tt.want[p].Fitness)
				}
				if !reflect.DeepEqual(got[p].DNA, tt.want[p].DNA) {
					t.Errorf("GenericStrategy.Select() DNA = %v, want %v", got[p].DNA, tt.want[p].DNA)
				}
			}
		})
	}
}

func TestGenericStrategy_Reproduce(t *testing.T) {
	rand.Seed(10)
	type fields struct {
		organismSize int
		alphabet     []byte
		fitnessFunc  FitnessFunc
	}
	type args struct {
		o1 Organism
		o2 Organism
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Organism
	}{
		{
			name: "Basic",
			fields: fields{
				organismSize: 10,
				alphabet:     GenerateAlphabet([]int{0, 1}),
				fitnessFunc:  func(o Organism) float64 { return 1.0 },
			},
			args: args{
				o1: Organism{
					DNA:     GenerateAlphabet([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}),
					Fitness: 0.0,
				},
				o2: Organism{
					DNA:     GenerateAlphabet([]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
					Fitness: -1.0,
				},
			},
			want: []Organism{
				{
					DNA:     GenerateAlphabet([]int{1, 1, 1, 1, 0, 0, 0, 0, 0, 0}),
					Fitness: 0.0,
				},
				{
					DNA:     GenerateAlphabet([]int{1, 1, 1, 1, 1, 1, 1, 1, 0, 0}),
					Fitness: 0.0,
				},
				{
					DNA:     GenerateAlphabet([]int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0}),
					Fitness: 0.0,
				},
			},
		},
	}
	// for _, tt := range tests {
	for i := 0; i < 3; i++ {
		tt := tests[0]
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				organismSize: tt.fields.organismSize,
				alphabet:     tt.fields.alphabet,
				fitnessFunc:  tt.fields.fitnessFunc,
			}
			got := g.Reproduce(tt.args.o1, tt.args.o2)

			if len(got.DNA) != tt.fields.organismSize {
				t.Errorf("GenericStrategy.Reproduce() size = %v, want %v", len(got.DNA), tt.fields.organismSize)
			}
			if !reflect.DeepEqual(got.DNA, tt.want[i].DNA) {
				t.Errorf("GenericStrategy.Reproduce() DNA = %v, want %v", got, tt.want[i])
			}
		})
	}
}

func TestGenericStrategy_Evolve(t *testing.T) {
	rand.Seed(999)
	type fields struct {
		organismSize       int
		populationSize     int
		selectionThreshold float64
		fitnessTarget      float64
		generationLimit    int
		mutationRate       float64
		alphabet           []byte
		fitnessFunc        FitnessFunc
	}
	tests := []struct {
		name         string
		fields       fields
		wantOrganism Organism
	}{
		{
			name: "xor solution",
			fields: fields{
				organismSize:       4,
				populationSize:     20,
				selectionThreshold: 0.9,
				fitnessTarget:      10.0,
				generationLimit:    100,
				mutationRate:       0.05,
				alphabet:           GenerateAlphabet([]int{0, 1}),
				fitnessFunc: func(o Organism) float64 {
					if len(o.DNA) != 4 {
						return 0.0
					}
					score := 0.0
					if o.DNA[0] == 0 {
						score++
					}
					if o.DNA[1] == 1 {
						score++
					}
					if o.DNA[2] == 1 {
						score++
					}
					if o.DNA[3] == 0 {
						score++
					}
					if reflect.DeepEqual(o.DNA, GenerateAlphabet([]int{0, 1, 1, 0})) {
						score = 10
					}
					return score
				},
			},
			wantOrganism: Organism{
				DNA:     GenerateAlphabet([]int{0, 1, 1, 0}),
				Fitness: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &GenericStrategy{
				organismSize:       tt.fields.organismSize,
				populationSize:     tt.fields.populationSize,
				selectionThreshold: tt.fields.selectionThreshold,
				fitnessTarget:      tt.fields.fitnessTarget,
				generationLimit:    tt.fields.generationLimit,
				mutationRate:       tt.fields.mutationRate,
				alphabet:           tt.fields.alphabet,
				fitnessFunc:        tt.fields.fitnessFunc,
			}
			gotOrganism, gotGenerations := g.Evolve()

			if !reflect.DeepEqual(gotOrganism, tt.wantOrganism) {
				t.Errorf("GenericStrategy.Evolve() got = %v, want %v", gotOrganism, tt.wantOrganism)
				if gotGenerations == tt.fields.generationLimit {
					t.Errorf("GenericStrategy.Evolve() reached max generations")
				}
			}
		})
	}
}
