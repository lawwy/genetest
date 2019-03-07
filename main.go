package main

import (
	"fmt"
	"genetest/utils"
)

type Task struct {
	Player *utils.Player
	Ground *utils.Ground
	Host   *utils.Host
}

func (t *Task) Run() {
	p := t.Player
	g := t.Ground
	h := t.Host
	fmt.Println("Genes:", p.Genes.GeneSeries())
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

func main() {
	p := utils.NewPlayer()
	g := utils.NewGround()
	h := &utils.Host{}
	p.Pos = g.RandomPos()
	t := &Task{
		Player: p,
		Ground: g,
		Host:   h,
	}
	t.Run()
}
