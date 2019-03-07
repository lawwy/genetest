package utils

import (
	"fmt"
	"testing"
)

func Test_RandomGene(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(RandomGene().GeneSeries())
	}
}

func Test_InitBoard(t *testing.T) {
	for i := 0; i < 3; i++ {
		b := InitBoard(9, 9, 10)
		for i := 0; i < len(b); i++ {
			fmt.Println(b[i])
		}
		fmt.Println("----------------------")
	}
	// fmt.Println(b)
}
