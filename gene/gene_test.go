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
	ff := []Fitness{
		Fitness{
			Item:  Gene{1, 2, 3, 4, 5},
			Score: 20.0,
		},
		Fitness{
			Item:  Gene{5, 4, 3, 2, 1},
			Score: 30.0,
		},
		Fitness{
			Item:  Gene{1, 2, 4, 3, 5},
			Score: 50.0,
		},
	}
	weights := InitWeights(ff)
	fmt.Println(weights)
}
