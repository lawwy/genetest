package utils

type Player struct {
	ID     int
	Pos    *Position
	Genes  GeneMap
	Score  int
	Ground *Ground
}

const STEPS int = 100

func NewPlayer(gm GeneMap) *Player {
	p := &Player{
		ID: AutoId(),
	}
	if gm != nil {
		p.Genes = gm
	} else {
		p.Genes = RandomGene()
	}
	return p
}

func (p *Player) NextMove(env *Env) move {
	key := ComputeGeneKey(env)
	if p.Genes[key] == 6 {
		return move(RandomInt(4))
	}
	return p.Genes[key]
}

func (p *Player) Ready() {
	p.Ground = NewGround()
	p.Pos = p.Ground.RandomPos()
	p.Score = 0
}

func (p *Player) Count(m move, success bool) {
	s := MoveScoreMap[m]
	if success {
		p.Score = p.Score + s.Success
	} else {
		p.Score = p.Score + s.Fail
	}
}

func (p *Player) Run() {
	g := p.Ground
	for i := 0; i < STEPS; i++ {
		env := g.GetAround(p.Pos)
		m := p.NextMove(env)
		success, np := g.Exec(p.Pos, m)
		p.Pos = np
		p.Count(m, success)
	}
}
