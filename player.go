package main

type Player struct {
	Pos   *Position
	Genes map[geneKey]move
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

type geneKey string

func NewPlayer() *Player {
	p := &Player{
		Pos: &Position{1, 1},
	}
	p.Genes = RandomGene()
	return p
}

func (p *Player) NextMove(env *Env) move {
	key := ComputeGeneKey(env)
	return p.Genes[key]
}
