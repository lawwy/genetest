package gene

import (
	"fmt"
)

var Rule = map[string]int{
	"PICK_JAR":     10,
	"PICK_NOTHING": -1,
	"HIT_WALL":     -5,
	"MOVE":         0,
	"STAND":        0,
}

type Gene []int

type Stage struct {
	Height  int
	Weight  int
	PR_jar  float64
	Board   [][]int
	Current *Position
}

type Position struct {
	X int
	Y int
}

func NewStage(h int, w int, pr float64) *Stage {
	stage := &Stage{
		Height: h,
		Weight: w,
		PR_jar: pr,
	}
	stage.initBoard()
	stage.RandomPosition()
	return stage
}

func (s *Stage) initBoard() {
	b := [][]int{}
	for i := 0; i < s.Height; i++ {
		r := make([]int, s.Weight)
		for j := 0; j < s.Weight; j++ {
			if RandomFloat() < s.PR_jar {
				r[j] = 1
			}
		}
		b = append(b, r)
	}
	s.Board = b
}

func (s *Stage) Print() {
	for i := 0; i < s.Height; i++ {
		for j := 0; j < s.Weight; j++ {
			switch s.GetState(j, i) {
			case 0:
				if s.Current.X == j && s.Current.Y == i {
					fmt.Print(" X")
					continue
				}
				fmt.Print(" o")
			case 1:
				if s.Current.X == j && s.Current.Y == i {
					fmt.Print(" @")
					continue
				}
				fmt.Print(" *")
			case 2:
				fmt.Print(" ?")
			}
		}
		fmt.Println()
	}
}

func Pow(x int, y int) int {
	p := 1
	for i := 0; i < y; i++ {
		p *= x
	}
	return p
}

func CountIndex(states []int) int {
	index := 0
	for i, s := range states {
		index += Pow(3, i) * s
	}
	return index
}

func (s *Stage) GetStates() []int {
	x := s.Current.X
	y := s.Current.Y
	up := s.GetState(x, y-1)
	right := s.GetState(x+1, y)
	down := s.GetState(x, y+1)
	left := s.GetState(x-1, y)
	mid := s.GetState(x, y)
	return append([]int{}, up, right, down, left, mid)
}

func (s *Stage) GetState(x int, y int) int {
	if x < 0 || y < 0 || x >= s.Weight || y >= s.Height {
		return 2
	}
	return s.Board[y][x]
}

func (s *Stage) RandomPosition() {
	s.Current = &Position{
		X: RandomInt(s.Weight),
		Y: RandomInt(s.Height),
	}
}

func (s *Stage) Move(act int) int {
	//执行动作
	//若成功
	//更改board状态，p移到下一个位置，并返回分数
	//若失败
	//返回分数
	p := s.Current
	switch act {
	case 0:
		//向上
		if s.GetState(p.X, p.Y-1) == 2 {
			return Rule["HIT_WALL"]
		}
		p.Y = p.Y - 1
		return Rule["MOVE"]
	case 1:
		//向右
		if s.GetState(p.X+1, p.Y) == 2 {
			return Rule["HIT_WALL"]
		}
		p.X = p.X + 1
		return Rule["MOVE"]
	case 2:
		//向下
		if s.GetState(p.X, p.Y+1) == 2 {
			return Rule["HIT_WALL"]
		}
		p.Y = p.Y + 1
		return Rule["MOVE"]
	case 3:
		//向左
		if s.GetState(p.X-1, p.Y) == 2 {
			return Rule["HIT_WALL"]
		}
		p.X = p.X - 1
		return Rule["MOVE"]
	case 4:
		//随机走
		act2 := RandomInt(4)
		return s.Move(act2)
	case 5:
		//捡箱
		if s.GetState(p.X, p.Y) == 1 {
			s.Board[p.Y][p.X] = 0
			return Rule["PICK_JAR"]
		}
		return Rule["PICK_NOTHING"]
	case 6:
		//不动
		return Rule["STAND"]
	}
	return 0
}

func Run(g Gene) int {
	stage := NewStage(10, 10, 0.4)
	moves := 200
	sum := 0
	for i := 0; i < moves; i++ {
		states := stage.GetStates()
		act := g[CountIndex(states)]
		sum += stage.Move(act)
	}
	return sum
}

func Exec(g Gene, ch chan<- *Fitness) {
	sum := 0
	times := 100
	for i := 0; i < times; i++ {
		sum += Run(g)
	}
	ch <- &Fitness{
		Item:  g,
		Score: float64(sum) / float64(times),
	}
}

func (g Gene) Copy() Gene {
	c := Gene{}
	for _, i := range g {
		c = append(c, i)
	}
	return c
}
