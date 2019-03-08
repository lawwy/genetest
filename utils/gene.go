package utils

import (
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func RandomInt(limit int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(limit)
}

const TOTAL int = 243

// type geneKey string
type geneKey int

/*
0:不动
1:向左
2:向右
3:向上
4:向下
5:拾起
*/
type move int

type GeneMap map[geneKey]move

func (gm GeneMap) GeneSeries() string {
	keys := []int{}
	for k := range gm {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	s := ""
	for _, k := range keys {
		s = s + strconv.Itoa(int(gm[geneKey(k)]))
	}
	return s
}

func CombineGenes(p, q GeneMap) GeneMap {
	n := GeneMap{}
	// mid := RandomInt(TOTAL - 1)
	// mid := 120

	for i := 0; i < TOTAL; i++ {
		if i%2 == 0 {
			n[geneKey(i)] = p[geneKey(i)]
		} else {
			n[geneKey(i)] = q[geneKey(i)]
		}
	}
	return n
}

func RandomGene() GeneMap {
	m := GeneMap{}
	// m := map[geneKey]move{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						k := computeGeneInt(q, p, k, j, i)
						m[geneKey(k)] = move(RandomInt(7))
					}
				}
			}
		}
	}
	return m
}

func computeGeneInt(up, down, left, right, mid int) int {
	return up*1 + down*3 + left*3*3 + right*3*3*3 + mid*3*3*3*3
}

func ComputeGeneKey(env *Env) geneKey {
	k := computeGeneInt(env.Up, env.Down, env.Left, env.Right, env.Mid)
	return geneKey(k)
}
