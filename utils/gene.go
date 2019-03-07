package utils

import (
	"math/rand"
	"time"
)

func RandomInt(limit int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(limit)
}

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

func RandomGene() map[geneKey]move {
	m := map[geneKey]move{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						k := q*1 + p*3 + k*3*3 + j*3*3*3 + i*3*3*3*3
						// k := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(p) + strconv.Itoa(q)
						m[geneKey(k)] = move(RandomInt(6))
					}
				}
			}
		}
	}
	return m
}

func ComputeGeneKey(env *Env) geneKey {
	// str := strconv.Itoa(env.Up) + strconv.Itoa(env.Down) + strconv.Itoa(env.Left) + strconv.Itoa(env.Right) + strconv.Itoa(env.Mid)
	k := env.Up*1 + env.Down*3 + env.Left*3*3 + env.Right*3*3*3 + env.Mid*3*3*3*3
	return geneKey(k)
}
