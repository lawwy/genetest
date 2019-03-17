package gene

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Gene []int

func (g Gene) Copy() Gene {
	c := Gene{}
	for _, i := range g {
		c = append(c, i)
	}
	return c
}

type ITask interface {
	Exec(Gene) float64
}

func RandomInt(limit int) int {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	return rand.Intn(limit)
}

func RandomFloat() float64 {
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	return rand.Float64()
}

func WriteGenes(path string, genes []Gene) error {
	f, err := os.Create(path)
	defer f.Close()
	if err != nil {
		return err
	}
	fmt.Println(len(genes))
	for _, g := range genes {
		ss := make([]string, len(g))
		for i, n := range g {
			ss[i] = strconv.Itoa(n)
		}
		_, err := f.WriteString(strings.Join(ss, ",") + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadPopulationFromFile(path string) ([]Gene, error) {
	gg := []Gene{}
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		ss := strings.Split(line, ",")
		g := Gene{}
		for _, s := range ss {
			n, _ := strconv.Atoi(s)
			g = append(g, n)
		}
		gg = append(gg, g)
	}
	return gg, nil
}
