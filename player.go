package main

import (
	"strconv"
)

type Player struct {
	Pos   *Position
	Genes map[string]move
}

/*
0:不动
1:向左
2:向右
3:向上
4:向下
5:拾起
*/
type move int

func NewPlayer() *Player {
	p := &Player{
		Pos: &Position{0, 0},
	}
	p.Genes = RandomGene()
	return p
}

func (p *Player) computeGeneKey(env *Env) string {
	return strconv.Itoa(env.Up) + strconv.Itoa(env.Down) + strconv.Itoa(env.Left) + strconv.Itoa(env.Right) + strconv.Itoa(env.Mid)
}

func (p *Player) NextMove(env *Env) move {
	key := p.computeGeneKey(env)
	return p.Genes[key]
}
