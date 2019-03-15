package gene

import (
	"fmt"
	"testing"
)

func Test_Random(t *testing.T) {
	fmt.Println(RandomInt(7))
	fmt.Println(RandomInt(7))
	fmt.Println(RandomFloat())
	fmt.Println(RandomFloat())
}

func Test_InitWeights(t *testing.T) {
	ff := []*Fitness{
		&Fitness{
			Item:  Gene{1, 2, 3, 4, 5},
			Score: 10.0,
		},
		&Fitness{
			Item:  Gene{5, 4, 3, 2, 1},
			Score: 30.0,
		},
		&Fitness{
			Item:  Gene{1, 2, 4, 3, 5},
			Score: 50.0,
		},
	}
	weights := InitWeights(ff)
	fmt.Println(weights)
}

func Test_GetInitPopulation(t *testing.T) {
	env := &Env{}
	env.PoplationSize = 3
	env.GeneSize = 5
	gg := env.GetInitPopulation()
	fmt.Println(gg)
}

// func Test_PopulationFitness(t *testing.T) {
// 	gg := []Gene{
// 		Gene{1, 2, 3, 4, 5},
// 		Gene{1, 2, 0, 4, 5},
// 		Gene{1, 2, 3, 8, 5},
// 	}
// 	ff := PopulationFitness(gg, func(Gene) float64 {
// 		return RandomFloat()
// 	})
// 	for _, f := range ff {
// 		fmt.Println(f)
// 	}
// }

func Test_Cross(t *testing.T) {
	f := Gene{0, 1, 2, 3, 4, 5, 6, 7, 8}
	m := Gene{8, 7, 6, 5, 4, 3, 2, 1, 0}
	g := Cross(f, m, 6)
	fmt.Println(g)
}

func Test_DoVariation(t *testing.T) {
	env := &Env{}
	env.MaxMutationCount = 1
	env.GeneSize = 6
	gene := Gene{0, 2, 3, 4, 5, 6}
	fmt.Println(env.DoVariation(gene))
}
