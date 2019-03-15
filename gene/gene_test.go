package gene

import (
	"fmt"
	"testing"
)

func Test_CountIndex(t *testing.T) {
	i1 := CountIndex([]int{0, 2, 0, 1, 2})
	i2 := CountIndex([]int{0, 2, 1, 1, 2})
	i3 := CountIndex([]int{1, 2, 2, 0, 0})
	if i1 != 195 || i2 != 204 || i3 != 25 {
		t.Fatal("fail")
	}
}

func Test_Stage(t *testing.T) {
	stage := NewStage(10, 10, 0.444344)
	stage.Print()
	fmt.Println("begin:", stage.GetStates(), stage.Current)
	s0 := stage.Move(0)
	fmt.Println("go up:", stage.GetStates(), stage.Current, s0)
	s1 := stage.Move(1)
	fmt.Println("go right:", stage.GetStates(), stage.Current, s1)
	s2 := stage.Move(2)
	fmt.Println("go down:", stage.GetStates(), stage.Current, s2)
	s3 := stage.Move(3)
	fmt.Println("go left:", stage.GetStates(), stage.Current, s3)
	s4 := stage.Move(4)
	fmt.Println("go random:", stage.GetStates(), stage.Current, s4)
	s5 := stage.Move(5)
	fmt.Println("pick jar:", stage.GetStates(), stage.Current, s5)
	s6 := stage.Move(6)
	fmt.Println("stand still:", stage.GetStates(), stage.Current, s6)
}

func Test_Run(t *testing.T) {
	env := &Env{}
	env.GeneSize = 243
	g := env.RandomGene()
	for i, s := range g {
		fmt.Println(i, ":", s)
	}
	score := Run(g)
	fmt.Println(score)
}

func Test_Show(t *testing.T) {
	fmt.Println("test")
	g := Gene{4, 0, 2, 1, 1, 2, 3, 4, 2, 4, 0, 1, 4, 1, 2, 3, 4, 2, 1, 0, 6, 0, 1, 4, 0, 0, 0, 3, 0, 2, 1, 4, 3, 4, 2, 4, 3, 3, 3, 1, 4, 1, 4, 3, 2, 0, 0, 0, 1, 4, 5, 3, 0, 5, 2, 1, 2, 4, 1, 4, 3, 2, 0, 1, 1, 1, 4, 4, 2, 6, 0, 0, 1, 1, 5, 1, 3, 6, 6, 2, 4, 5, 5, 5, 5, 5, 1, 5, 5, 2, 2, 5, 4, 5, 5, 2, 5, 0, 3, 5, 5, 4, 1, 1, 1, 5, 5, 3, 5, 5, 5, 5, 5, 1, 3, 5, 5, 2, 4, 5, 2, 5, 2, 3, 5, 5, 5, 5, 0, 1, 1, 0, 5, 0, 6, 5, 5, 2, 5, 5, 2, 3, 2, 0, 5, 5, 5, 2, 5, 6, 1, 5, 1, 5, 4, 0, 1, 1, 3, 4, 6, 5, 4, 1, 4, 0, 0, 5, 2, 1, 5, 3, 5, 0, 2, 2, 3, 2, 2, 1, 4, 2, 6, 4, 6, 1, 5, 0, 0, 0, 3, 2, 4, 3, 3, 2, 4, 5, 3, 5, 5, 2, 6, 1, 4, 6, 3, 1, 0, 5, 3, 3, 5, 6, 0, 2, 2, 1, 6, 1, 3, 3, 5, 6, 6, 5, 2, 3, 0, 1, 1, 4, 1, 5, 2, 3, 0, 2, 1, 3, 5, 2, 4}
	Show(g)
}

// func Test_Exec(t *testing.T) {
// 	env := &Env{}
// 	env.GeneSize = 243
// 	g := env.RandomGene()
// 	score := Exec(g)
// 	fmt.Println(score)
// }
