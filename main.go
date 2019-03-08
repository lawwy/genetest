package main

import (
	"fmt"
	"genetest/utils"
	"sort"
)

const GENS int = 100

func main() {
	pp := RandomPlayers()
	for i := 0; i < GENS; i++ {
		for _, p := range pp {
			p.Ready()
			p.Run()
		}
		SortPlayers(pp)
		pp = MergeAndSelect(pp)
	}
	for _, p := range pp {
		fmt.Println(p.Score)
	}
}

type PlayersSorter []*utils.Player

func (pp PlayersSorter) Len() int {
	return len(pp)
}

func (pp PlayersSorter) Swap(i, j int) {
	pp[i], pp[j] = pp[j], pp[i]
}

func (pp PlayersSorter) Less(i, j int) bool {
	return pp[i].Score > pp[j].Score
}

func SortPlayers(pp []*utils.Player) {
	ps := PlayersSorter(pp)
	sort.Sort(ps)
}

func RandomPlayers() []*utils.Player {
	pp := []*utils.Player{}
	for i := 0; i < 100; i++ {
		p := utils.NewPlayer(nil)
		pp = append(pp, p)
	}
	return pp
}

func MergeAndSelect(pp []*utils.Player) []*utils.Player {
	m := map[int]*utils.Player{}
	for i, p := range pp {
		m[i] = p
	}
	npp := []*utils.Player{}
	for i := 0; i < 20; i++ {
		g := utils.CombineGenes(pp[i].Genes, pp[39-i].Genes)
		np := utils.NewPlayer(g)
		npp = append(npp, np)
	}
	pp = pp[:80]
	pp = append(pp, npp...)
	return pp
}
