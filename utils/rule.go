package utils

type Cond struct {
	Success int
	Fail    int
}

var MoveScoreMap map[move]*Cond = map[move]*Cond{
	0: &Cond{1, -1},  //向左
	1: &Cond{1, -1},  //向右
	2: &Cond{1, -1},  //向上
	3: &Cond{1, -1},  //向下
	4: &Cond{2, 0},   //拾起
	5: &Cond{-1, -1}, //不动
	// 6: &Cond{2, -1},
}
