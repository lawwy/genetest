package utils

type Position struct {
	X int
	Y int
}

func (p *Position) Left() *Position {
	return &Position{p.X - 1, p.Y}
}

func (p *Position) Right() *Position {
	return &Position{p.X + 1, p.Y}
}

func (p *Position) Up() *Position {
	return &Position{p.X, p.Y - 1}
}

func (p *Position) Down() *Position {
	return &Position{p.X, p.Y + 1}
}
