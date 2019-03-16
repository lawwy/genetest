package main

import (
	"genetest/gene"
)

const (
	//种群个体数量
	POPULATION_SIZE       = 200 //200
	VARIATION_PROBABILITY = 0.078
	EVAL_TIMES            = 10
	CROSS_PROBABILITY     = 0.82
	MAX_VARI_COUNT        = 10
)

func main() {
	env := &gene.Env{
		EvalTime:         EVAL_TIMES,
		MaxMutationCount: MAX_VARI_COUNT,
		PR_Cross:         CROSS_PROBABILITY,
		PR_Mutation:      VARIATION_PROBABILITY,
		PoplationSize:    POPULATION_SIZE,
	}

	env.Task = &gene.RobotTask{}
	env.GeneSize = gene.ROBOT_GENE_SIZE
	env.GeneRange = gene.ROBOT_GENE_RANGE

	population, err := gene.ReadPopulationFromFile("./robot.gene")
	checkErr(err)
	population = env.Start(population)
	err = gene.WriteGenes("./robot.gene", population)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
