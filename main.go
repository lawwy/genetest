package main

import (
	"genetest/gene"
)

const (
	//种群个体数量
	POPULATION_SIZE       = 200 //200
	TRY_TIMES             = 100
	MOVE_TIMES            = 200
	JAR_PROBABILITY       = 0.5
	VARIATION_PROBABILITY = 0.078
	EVAL_TIMES            = 2000
	CROSS_PROBABILITY     = 0.82
	MAX_VARI_COUNT        = 10
)

const GENE_SIZE = 243 //243

var Rule = map[string]int{
	"PICK_JAR":     10,
	"PICK_NOTHING": -1,
	"HIT_WALL":     -5,
}

func main() {
	env := &gene.Env{}
	env.GeneSize = GENE_SIZE
	env.EvalTime = EVAL_TIMES
	env.MaxMutationCount = MAX_VARI_COUNT
	env.MoveTime = MOVE_TIMES
	env.PR_Cross = CROSS_PROBABILITY
	env.PR_Jar = JAR_PROBABILITY
	env.PR_Mutation = VARIATION_PROBABILITY
	env.PoplationSize = POPULATION_SIZE
	env.Rule = Rule
	env.TryTime = TRY_TIMES
	env.Start()
}
