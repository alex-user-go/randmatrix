package randmatrix

import (
	"errors"
	"math/rand"
	"time"
)

func CreateArray(x, y int, min int, max int) ([][]int, error) {
	gen := NewIntGenerator(time.Now().UnixNano())
	matrix := make([][]int, x)
	for i := 0; i < x; i++ {
		line := make([]int, y)
		for j := 0; j < y; j++ {
			v, err := gen.Unique(min, max)
			if err != nil {
				return nil, err
			}
			line[j] = v
		}
		matrix[i] = line
	}
	return matrix, nil
}

type IntGenerator struct {
	seed       int64
	uniqueVals map[int]bool
}

func NewIntGenerator(seed int64) *IntGenerator {
	g := new(IntGenerator)
	rand.Seed(g.seed)
	g.uniqueVals = map[int]bool{}
	return g
}

var (
	ErrInvalidRange  = errors.New("invalid use of min and max value")
	ErrRangeTooSmall = errors.New("not enough numbers in a range")
)

func (g *IntGenerator) Rand(min, max int) (int, error) {

	if min >= max {
		return 0, ErrInvalidRange
	}
	randNum := rand.Intn(max-min) + min
	return randNum, nil
}

func (g *IntGenerator) Unique(min, max int) (int, error) {
	for {
		v, err := g.Rand(min, max)
		if err != nil {
			return 0, err
		}
		if _, exist := g.uniqueVals[v]; !exist {
			g.uniqueVals[v] = true
			return v, nil
		}
		//no more numbers
		if len(g.uniqueVals) == (max - min) {
			return 0, ErrRangeTooSmall
		}
	}
}
