package main

import (
	"fmt"
	"genetest/utils"
)

//TODO:主进化逻辑
func Run(ts []*utils.Player) []*utils.Player {
	r := make(chan *Res,100)
	for _,p := range ts {
		go Play(p,r)
	}
	end := make(chan bool)
	rr := []*Res{}
	for c := range r {
		rr = append(rr,c)
		if len(rr) == 100 {
			end <- true
		}
	}
	<-end
	return MergeAndSelect(rr)
}

func MergeAndSelect(rr []*Res) []*utils.Player {
	//TODO
	ps := []*utils.Player{}
	for _,r := range rr {
		ps = append(ps,r.Player)
	}
	return ps
}

type Res struct {
	Player *utils.Player
	Score int
}

func Play(p *utils.Player,res chan *Res){
	g := utils.NewGround()
	h := &utils.Host{}
	p.Pos = g.RandomPos()
	for i := 0; i < 1000; i++ {
		// fmt.Println("pos:", p.Pos)
		env := g.GetAround(p.Pos)
		// fmt.Println("env:", env)
		m := p.NextMove(env)
		// fmt.Println("m:", m)
		success, np := g.Exec(p.Pos, m)
		// fmt.Println("success:", success)
		// fmt.Println("new Pos:", np)
		p.Pos = np
		h.Count(m, success)
	}
	// fmt.Println(h.Score)
	res <- &Res{
		Player:p,
		Score:h.Score,
	}
}

func main() {
	//TODO:test
	fmt.Println("Start")
	ps := []*utils.Player{}
	for i:=0;i<100;i++{
		ps = append(ps,utils.NewPlayer())
	}
	for g:=0;g<100;g++{
		ps = Run(ps)
	}
	//TODO:输出结果
}
