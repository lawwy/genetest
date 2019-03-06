package main

import (
	"math/rand"
	"strconv"
)

func RandomGene() map[string]move {
	m := map[string]move{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for p := 0; p < 3; p++ {
					for q := 0; q < 3; q++ {
						k := strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k) + strconv.Itoa(p) + strconv.Itoa(q)
						m[k] = move(rand.Intn(6))
					}
				}
			}
		}
	}
	return m
}
