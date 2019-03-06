package main

import (
	"math/rand"
	"strconv"
)

func RandomGene() map[geneKey]move {
	m := map[geneKey]move{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						k := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(p) + strconv.Itoa(q)
						m[geneKey(k)] = move(rand.Intn(6))
					}
				}
			}
		}
	}
	return m
}

func ComputeGeneKey(env *Env) geneKey {
	str := strconv.Itoa(env.Up) + strconv.Itoa(env.Down) + strconv.Itoa(env.Left) + strconv.Itoa(env.Right) + strconv.Itoa(env.Mid)
	return geneKey(str)
}