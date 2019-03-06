package main

import (
	"fmt"
	"genetest/utils"
)

func main() {
	p := utils.NewPlayer()
	g := utils.NewGround()
	h := &utils.Host{}
	for i:=0;i<1000;i++{
		env := g.GetAround(p.Pos)
		m := p.NextMove(env)
		success,np := g.Exec(p.Pos,m)
		p.Pos = np
		h.Count(m,success)
	}
	//BUG:ALWAYS BE 0
	fmt.Println(h.Score)
}
