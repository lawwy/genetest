package gene

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

type Env struct {
	PoplationSize    int
	TryTime          int
	MoveTime         int
	PR_Jar           float64
	PR_Mutation      float64
	PR_Cross         float64
	EvalTime         int
	MaxMutationCount int
	GeneSize         int
	GeneRange        int
	Rule             map[string]int
	RecordPath       string
	Task             ITask
}

func init() {
	rand.Seed(time.Now().Unix())
	fmt.Println("random")
}

func (env *Env) Start(population []Gene) []Gene {
	fmt.Println("Start:")
	if len(population) < env.PoplationSize {
		population = env.GetInitPopulation()
	}
	fmt.Println("Origin Population:", len(population))
	for i := 0; i < env.EvalTime; i++ {
		fmt.Println("GEN:", i)
		scoreList := env.PopulationFitness(population)
		population = env.Evolve(scoreList)
	}
	return population
}

func (env *Env) GetInitPopulation() []Gene {
	gg := []Gene{}
	for i := 0; i < env.PoplationSize; i++ {
		g := env.RandomGene()
		gg = append(gg, g)
	}
	return gg
}

//TODO:较由各自的任务生成随机基因？
func (env *Env) RandomGene() Gene {
	g := Gene{}
	for j := 0; j < env.GeneSize; j++ {
		n := RandomInt(env.GeneRange)
		g = append(g, n)
	}
	return g
}

func (env *Env) PopulationFitness(gg []Gene) []*Fitness {
	ch := make(chan *Fitness, len(gg))
	ff := []*Fitness{}
	for _, g := range gg {
		go env.Work(g, ch)
	}
	for i := 0; i < len(gg); i++ {
		r := <-ch
		ff = append(ff, r)
	}
	return ff
}

func (env *Env) Work(g Gene, res chan *Fitness) {
	score := env.Task.Exec(g)
	res <- &Fitness{
		Item:  g,
		Score: score,
	}
}

func (env *Env) Evolve(ff []*Fitness) []Gene {
	fmt.Println("Evolve")
	sort.Sort(FitnessSlice(ff))
	fmt.Println("平均分", FitnessSlice(ff).Average())
	fmt.Println("最高分：", ff[len(ff)-1].Score)
	fmt.Println(ff[len(ff)-1].Item)
	getIndex := GetWeightIndexFunc(InitWeights(ff))
	newPopulation := []Gene{}
	for len(newPopulation) < env.PoplationSize {
		father := ff[getIndex(RandomFloat())].Item.Copy()
		mother := ff[getIndex(RandomFloat())].Item.Copy()
		var child1, child2 Gene
		if crossRate := RandomFloat(); crossRate < env.PR_Cross {
			rPos := env.CrossPoint()
			child1 = Cross(father, mother, rPos)
			child2 = Cross(mother, father, rPos)
		} else {
			child1 = father
			child2 = mother
		}
		if r1 := RandomFloat(); r1 < env.PR_Mutation {
			child1 = env.DoVariation(child1)
		}
		if r2 := RandomFloat(); r2 < env.PR_Mutation {
			child2 = env.DoVariation(child2)
		}
		newPopulation = append(newPopulation, child1, child2)
	}
	return newPopulation
}

func Cross(father Gene, mother Gene, rPos int) Gene {
	g := Gene{}
	g = append(g, father[:rPos]...)
	g = append(g, mother[rPos:]...)
	return g
}

func GetWeightIndexFunc(weights []float64) func(float64) int {
	return func(r float64) int {
		if r > 1 {
			panic("r should be less than 1")
		}
		for i := 0; i < len(weights); i++ {
			if r <= weights[i] {
				return i
			}
		}
		return -1
	}
}

func InitWeights(ff []*Fitness) []float64 {
	var total, adjust float64
	if ff[0].Score < 0 {
		adjust = math.Abs(ff[0].Score) + 1
	}
	weights := []float64{}
	for _, f := range ff {
		f.Score = math.Pow(f.Score+adjust, 2) //可调整
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

func (env *Env) CrossPoint() int {
	//不能是0或者GeneSize
	return RandomInt(env.GeneSize-1) + 1
}

func (env *Env) DoVariation(g Gene) Gene {
	count := RandomInt(env.MaxMutationCount) + 1
	for i := 0; i < count; i++ {
		rIndex := RandomInt(env.GeneSize)
		ng := RandomInt(env.GeneRange)
		for g[rIndex] == ng {
			ng = RandomInt(env.GeneRange)
		}
		g[rIndex] = ng
	}
	return g
}

type Fitness struct {
	Item  Gene
	Score float64
}

type FitnessSlice []*Fitness

func (ff FitnessSlice) Len() int {
	return len(ff)
}

func (ff FitnessSlice) Swap(i, j int) {
	ff[i], ff[j] = ff[j], ff[i]
}

func (ff FitnessSlice) Less(i, j int) bool {
	return ff[i].Score < ff[j].Score
}

func (ff FitnessSlice) Average() float64 {
	sum := 0.0
	for _, f := range ff {
		sum += f.Score
	}
	return sum / float64(len(ff))
}
