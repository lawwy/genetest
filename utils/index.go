package utils

import (
	"fmt"
	"sort"
)

const GENS int = 100

func Start() {
	pp := RandomPlayers()
	for i := 0; i < GENS; i++ {
		for _, p := range pp {
			p.Ready()
			p.Run()
		}
		SortPlayers(pp)
		pp = MergeAndSelect(pp)
	}
	fmt.Println("LAST:")
	for _, p := range pp {
		fmt.Print(p.Score, " ")
	}
}

type PlayersSorter []*Player

func (pp PlayersSorter) Len() int {
	return len(pp)
}

func (pp PlayersSorter) Swap(i, j int) {
	pp[i], pp[j] = pp[j], pp[i]
}

func (pp PlayersSorter) Less(i, j int) bool {
	return pp[i].Score > pp[j].Score
}

func SortPlayers(pp []*Player) {
	ps := PlayersSorter(pp)
	sort.Sort(ps)
}

func RandomPlayers() []*Player {
	pp := []*Player{}
	for i := 0; i < 100; i++ {
		p := NewPlayer(nil)
		// fmt.Println(p.Genes.GeneSeries())
		pp = append(pp, p)
	}
	return pp
}

func MergeAndSelect(pp []*Player) []*Player {
	m := map[int]*Player{}
	for i, p := range pp {
		m[i] = p
	}
	npp := []*Player{}
	for i := 0; i < 20; i++ {
		g := CombineGenes(pp[i].Genes, pp[39-i].Genes)
		np := NewPlayer(g)
		npp = append(npp, np)
	}
	pp = pp[:80]
	pp = append(pp, npp...)
	return pp
}
