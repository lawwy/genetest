package main

import (
	"fmt"
)

func main() {
	p := NewPlayer()
	g := NewGround()
	h := &Host{}
	env := g.GetAround(p.Pos)
	m := p.NextMove(env)
	success,np := g.Exec(p.Pos,m)
	p.Pos = np
	h.Count(m,success)
	fmt.Println(h.Score)
}
