package main

import (
	"genetest/gene"
)

const (
	//种群个体数量
	POPULATION_SIZE       = 100 //200
	VARIATION_PROBABILITY = 0.078
	EVAL_TIMES            = 1000
	CROSS_PROBABILITY     = 0.82
	MAX_VARI_COUNT        = 8
	// MAX_VARI_COUNT        = 10

)

func main() {
	env := &gene.Env{
		EvalTime:         EVAL_TIMES,
		MaxMutationCount: MAX_VARI_COUNT,
		PR_Cross:         CROSS_PROBABILITY,
		PR_Mutation:      VARIATION_PROBABILITY,
		PoplationSize:    POPULATION_SIZE,
	}
	chooseTask("cell_classify", env)

	// population, err := gene.ReadPopulationFromFile("./robot.gene")
	// checkErr(err)
	// population = env.Start(population)
	env.Start(nil)
	// err = gene.WriteGenes("./robot.gene", population)
	// checkErr(err)
}

func chooseTask(t string, env *gene.Env) {
	switch t {
	case "clean_robot":
		env.Task = &gene.RobotTask{}
		env.GeneSize = gene.ROBOT_GENE_SIZE
		env.GeneRange = gene.ROBOT_GENE_RANGE
	case "cell_classify":
		env.Task = &gene.CellTask{}
		env.GeneSize = gene.CELL_GENE_SIZE
		env.GeneRange = gene.CELL_GENE_RANGE
	}
	return
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
