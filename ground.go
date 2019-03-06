package main

type Env struct {
	Up    int
	Down  int
	Left  int
	Right int
	Mid   int
}

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

type Ground struct {
	Board [][]int
}

//0表示墙，1表示空地，2表示食物

func (g *Ground) GetAround(p *Position) *Env {
	return &Env{
		Up:    g.Value(&Position{p.X, p.Y - 1}),
		Down:  g.Value(&Position{p.X, p.Y + 1}),
		Left:  g.Value(&Position{p.X - 1, p.Y}),
		Right: g.Value(&Position{p.X + 1, p.Y}),
		Mid:   g.Value(p),
	}
}

func (g *Ground) Exec(p *Position, m move) (bool, *Position) {
	//TODO:这个方法好像有点复杂
	switch m {
	case 0:
		return true, p
		// case 1:
		// 	if g.IsWall(p.Left())
		// 		return false
		// case 2:
		// case 3:
		// case 4:
		// case 5:
	}
	return true, p
}

func (g *Ground) IsWall(p *Position) bool {
	if g.Board[p.X][p.Y] == 0 {
		return true
	}
	return false
}

func (g *Ground) Value(p *Position) int {
	return g.Board[p.X][p.Y]
}

func (g *Ground) Init() {

}
