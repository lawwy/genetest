package utils

//QUESTION:player需要有更多的职能，计分，执行步骤等
type Player struct {
	ID int
	Pos   *Position
	Genes GeneMap
}

func NewPlayer() *Player {
	p := &Player{
		ID: AutoId(),		
		Pos: &Position{1, 1},
	}
	p.Genes = RandomGene()
	return p
}

func (p *Player) NextMove(env *Env) move {
	key := ComputeGeneKey(env)
	return p.Genes[key]
}
