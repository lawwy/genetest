package main

import (
	"fmt"
	"testing"
)

func Test_RandomGene(t *testing.T){
	gs := RandomGene()
	fmt.Println(gs)
}

func Test_InitBoard(t *testing.T){
	b := InitBoard(9,9,10)
	fmt.Println(b)
}