package gene

import (
	"fmt"
	"time"
)

const (
	CELL_GENE_SIZE  = 128
	CELL_GENE_RANGE = 2
	CELL_MOVE_TIMES = 60
	CELL_TRY_TIMES  = 20
	GROUND_WEIGHT   = 101
	RADIUS          = 3
)

type Ground struct {
	Weight   int
	Row      []int
	Answer   int
	PR_Black float64
	Radius   int
}

func NewGround(w int, pr float64, radius int) *Ground {
	g := &Ground{}
	g.Weight = w
	g.PR_Black = pr
	g.Radius = radius
	g.init()
	return g
}

func (s *Ground) init() {
	s.Row = make([]int, s.Weight)
	blacks := 0
	for i, _ := range s.Row {
		if RandomFloat() < s.PR_Black {
			s.Row[i] = 1
			blacks += 1
		}
	}
	s.Answer = 0
	if blacks > s.Weight/2 {
		s.Answer = 1
	}
}

func (s *Ground) getStates(x int) []int {
	// ss := []int{x - 3, x - 2, x - 1, x, x + 1, x + 2, x + 3}
	ss := []int{}
	l := len(s.Row)
	for i := -s.Radius; i <= s.Radius; i++ {
		if i+x < 0 {
			ss = append(ss, s.Row[l+i+x])
			continue
		}
		if i+x > l-1 {
			ss = append(ss, s.Row[i+x-l])
			continue
		}
		ss = append(ss, s.Row[i+x])
	}
	return ss
}

func (s *Ground) CountIndex(states []int) int {
	index := 0
	for i, s := range states {
		index += Pow(2, i) * s
	}
	return index
}

func (s *Ground) Move(m int, i int) int {
	switch m {
	case 0:
		//变白
		return 0
	case 1:
		//变黑
		return 1
		// case 2:
		// 	return s.Move(RandomInt(2), i)
		// case 3:
		// 	return s.Row[i]
		//不变
	}
	return 0
}

func (s *Ground) Wave(g Gene) {
	//TODO
	r := make([]int, s.Weight)
	for i, _ := range s.Row {
		states := s.getStates(i)
		move := g[s.CountIndex(states)]
		r[i] = s.Move(move, i)
	}
	s.Row = r
	// return r
}

func (s *Ground) Score() float64 {
	score := 0
	for _, r := range s.Row {
		if r == s.Answer {
			score += 1
		}
	}
	return float64(score)
}

func (s *Ground) Print() {
	for _, n := range s.Row {
		if n == 0 {
			fmt.Print("o")
		} else {
			fmt.Print("·")
		}
	}
	fmt.Println()
}

func (s *Ground) Show(g Gene, times int) {
	for i := 0; i < times; i++ {
		s.Wave(g)
		s.Print()
		time.Sleep(time.Second)
	}
}

func RunCellTask(g Gene) float64 {
	s := NewGround(GROUND_WEIGHT, RandomFloat(), RADIUS)
	for i := 0; i < CELL_MOVE_TIMES; i++ {
		s.Wave(g)
	}
	return s.Score()
}

type CellTask struct {
}

func (c *CellTask) Exec(g Gene) float64 {
	sum := 0.0
	for i := 0; i < CELL_TRY_TIMES; i++ {
		sum += RunCellTask(g)
	}
	return sum / float64(CELL_TRY_TIMES)
}
