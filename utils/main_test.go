package utils

import (
	"fmt"
	"testing"
)

//BUG:并不随机
func Test_RandomGene(t *testing.T) {
	gs := RandomGene()
	for k, v := range gs {
		fmt.Println(k, v)
	}
	// fmt.Println(gs)
}

func Test_InitBoard(t *testing.T) {
	b := InitBoard(9, 9, 10)
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	// fmt.Println(b)
}
