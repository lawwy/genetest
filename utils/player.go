package utils

type Player struct {
	Pos   *Position
	Genes GeneMap
}

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
