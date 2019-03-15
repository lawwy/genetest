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
	stage := NewStage(10, 10, 0.4344)
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

// func Test_Exec(t *testing.T) {
// 	env := &Env{}
// 	env.GeneSize = 243
// 	g := env.RandomGene()
// 	score := Exec(g)
// 	fmt.Println(score)
// }
