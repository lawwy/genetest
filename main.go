package main

import (
	"fmt"
	"genetest/utils"
)

func main() {
	p := utils.NewPlayer()
	g := utils.NewGround()
	h := &utils.Host{}
	for i := 0; i < len(g.Board); i++ {
		fmt.Println(g.Board[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Println("pos:", p.Pos)
		env := g.GetAround(p.Pos)
		fmt.Println("env:", env)
		m := p.NextMove(env)
		fmt.Println("m:", m)
		success, np := g.Exec(p.Pos, m)
		fmt.Println("success:", success)
		fmt.Println("new Pos:", np)
		p.Pos = np
		h.Count(m, success)
	}
	fmt.Println(h.Score)
}
