package main

import (
	"math/rand"
)

type Env struct {
	Up    int
	Down  int
	Left  int
	Right int
	Mid   int
}

type Ground struct {
	Board [][]int
	Height int
	Weight int
	Solids int
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
	switch m {
	case 0:
		//不动
		return true, p
	case 1:
		//向左
		left := p.Left();
		if g.IsWall(left) {
			return false,p
		}
		return true,left
	case 2:
		right := p.Right();
		if g.IsWall(right) {
			return false,p
		}
		return true,right
	case 3:
		up := p.Up();
		if g.IsWall(up) {
			return false,p
		}
		return true,up
	case 4:
		down := p.Down();
		if g.IsWall(down) {
			return false,p
		}
		return true,down
	case 5:
		if g.IsSolid(p){
			g.Board[p.X][p.Y] = 1
			return true,p
		}
		return false,p
	}
	return true, p
}

func (g *Ground) IsSolid(p *Position) bool {
	if g.Board[p.X][p.Y] == 2 {
		return true
	}
	return false
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

//TODO:test
func InitBoard(h int,w int,s int) [][]int {
	b := [][]int{}
	b = append(b,FillSlice(w+2,0))
	for i:=0;i<h;i++{
		r := FillSlice(w,1)
		r = append([]int{0},r...)
		r = append(r,0)
		b = append(b,r)
	}
	b = append(b,FillSlice(w+2,0))
	for i:=0;i<s;i++{
		p := rand.Intn(h-1)+1
		q := rand.Intn(w-1)+1
		b[p][q] = 2
	}
	return b
}

func NewGround() *Ground {
	g := &Ground{
		Height:9,
		Weight:9,
		Solids:10,
	}
	g.Board = InitBoard(g.Height,g.Weight,g.Solids)
	return g
}

func FillSlice(l int,v int) []int {
	s := []int{}
	for i:=0;i<l;i++{
		s = append(s,v)
	}
	return s
}