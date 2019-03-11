package gene

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	//种群个体数量
	POPULATION_SIZE       = 200
	TRY_TIMES             = 100
	MOVE_TIMES            = 200
	JAR_PROBABILITY       = 0.5
	VARIATION_PROBABILITY = 0.078
	EVAL_TIMES            = 2000
	CROSS_PROBABILITY     = 0.82
	MAX_VARI_COUNT        = 10
)

const GENE_SIZE = 243

var Rule = map[string]int{
	"PICK_JAR":     10,
	"PICK_NOTHING": -1,
	"HIT_WALL":     -5,
}

func Start() {
	population := GetInitPopulation()
	for i := 0; i < EVAL_TIMES; i++ {
		scoreList := PopulationFitness(population)
		population = Evolve(scoreList)
	}
}

type Gene []int
type Fitness struct {
	Item  Gene
	Score float64
}

type FitnessSlice []Fitness

func (ff FitnessSlice) Len() int {
	return len(ff)
}

func (ff FitnessSlice) Swap(i, j int) {
	ff[i], ff[j] = ff[j], ff[i]
}

func (ff FitnessSlice) Less(i, j int) bool {
	return ff[i].Score < ff[j].Score
}

func RandomInt(limit int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(limit)
}

func RandomFloat() float64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Float64()
}

func (g *Gene) Exec() float64 {
	//TODO
	return RandomFloat()
}

func GetInitPopulation() []Gene {
	gg := []Gene{}
	for i := 0; i < POPULATION_SIZE; i++ {
		g := Gene{}
		for j := 0; j < GENE_SIZE; j++ {
			n := RandomInt(7)
			g = append(g, n)
		}
		gg = append(gg, g)
	}
	return gg
}

func PopulationFitness(gg []Gene) []Fitness {
	ff := []Fitness{}
	for _, g := range gg {
		score := g.Exec()
		ff = append(ff, Fitness{g, score})
	}
	return ff
}

func Evolve(ff []Fitness) []Gene {
	sort.Sort(FitnessSlice(ff))
	fmt.Println("最高分")
	fmt.Println(ff[len(ff)-1].Score)
	fmt.Println(ff[len(ff)-1].Item)
	getIndex := GetWeightIndexFunc(InitWeights(ff))
	newPopulation := []Gene{}
	for len(newPopulation) <= GENE_SIZE {
		father := Gene{}
		mother := Gene{}
		copy(father, ff[getIndex(RandomFloat())].Item)
		copy(mother, ff[getIndex(RandomFloat())].Item)
		var child1, child2 Gene
		if crossRate := RandomFloat(); crossRate < CROSS_PROBABILITY {
			rPos := CrossPoint()
			child1 = Cross(father, mother, rPos)
			child2 = Cross(mother, father, rPos)
		} else {
			child1 = father
			child2 = mother
		}
		if r1 := RandomFloat(); r1 < VARIATION_PROBABILITY {
			child1 = DoVariation(child1)
		}
		if r2 := RandomFloat(); r2 < VARIATION_PROBABILITY {
			child2 = DoVariation(child2)
		}
		newPopulation = append(newPopulation, child1, child2)
	}
	return []Gene{}
}

func Cross(father Gene, mother Gene, rPos int) Gene {
	g := Gene{}
	g = append(g, father[:rPos]...)
	g = append(g, mother[rPos:]...)
	return g
}

func GetWeightIndexFunc(weights []float64) func(float64) int {
	return func(r float64) int {
		for i := 0; i < len(weights); i++ {
			if r <= weights[i] {
				return i
			}
		}
		return 0
	}
}

func InitWeights(ff []Fitness) []float64 {
	adjust := math.Abs(ff[0].Score) + 1
	total := 0.0
	weights := []float64{}
	for _, f := range ff {
		f.Score = f.Score + adjust //可调整
		total = total + f.Score
	}
	prev := 0.0
	for _, f := range ff {
		w := f.Score/total + prev
		weights = append(weights, w)
		prev = w
	}
	return weights
}

func CrossPoint() int {
	//TODO
	return RandomInt(GENE_SIZE)
}

func DoVariation(g Gene) Gene {
	count := RandomInt(MAX_VARI_COUNT) + 1
	for i := 0; i < count; i++ {
		rIndex := RandomInt(GENE_SIZE)
		ng := RandomInt(7)
		for g[rIndex] == ng {
			ng = RandomInt(7)
		}
		g[rIndex] = ng
	}
	return g
}
