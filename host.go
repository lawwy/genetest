package main

type Cond struct {
	Success int
	Fail    int
}

var MoveScoreMap map[move]*Cond = map[move]*Cond{
	0: &Cond{0, -1},
	1: &Cond{0, -1},
	2: &Cond{0, -1},
	3: &Cond{0, -1},
	4: &Cond{0, -1},
	5: &Cond{0, -1},
	6: &Cond{2, -1},
}

type Host struct {
	Score int
}

func (h *Host) Count(m move, success bool) {
	s := MoveScoreMap[m]
	if success {
		h.Score = h.Score + s.Success
	} else {
		h.Score = h.Score + s.Fail
	}
}
