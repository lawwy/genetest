package utils

type Cond struct {
	Success int
	Fail    int
}

var MoveScoreMap map[move]*Cond = map[move]*Cond{
	0: &Cond{0, -1},
	1: &Cond{1, -1},
	2: &Cond{1, -1},
	3: &Cond{1, -1},
	4: &Cond{1, -1},
	5: &Cond{2, -1},
	// 6: &Cond{2, -1},
}
