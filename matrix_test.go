package randmatrix

import (
	"testing"
	"time"
)

func TestRand(t *testing.T) {
	tests := []struct {
		min int
		max int
		err error
	}{
		{min: 1, max: 5, err: nil},
		{min: 2, max: 2, err: ErrInvalidRange},
		{min: 5, max: 1, err: ErrInvalidRange},
		{min: -100, max: -2, err: nil},
		{min: -100, max: -102, err: ErrInvalidRange},
	}
	gen := NewIntGenerator(time.Now().UnixNano())
	for _, v := range tests {
		res, err := gen.Rand(v.min, v.max)
		if err != nil {
			if err != v.err {
				t.Fatalf("expected error: %v, got: %v", v.err, err)
			}
		} else {
			if res > v.max || res < v.min {
				t.Fatalf("want a number in a range: %d - %d, got: %d", v.min, v.max, res)
			}
		}
	}
}
func TestUnique(t *testing.T) {
	tests := []struct {
		min      int
		max      int
		attempts int
		err      error
	}{
		{min: 0, max: 9, attempts: 9, err: nil},
		{min: 1, max: 5, attempts: 100, err: ErrRangeTooSmall},
		{min: 2, max: 2, attempts: 2, err: ErrInvalidRange},
		{min: 5, max: 1, attempts: 2, err: ErrInvalidRange},
		{min: 1, max: 2, attempts: 1, err: nil},
		{min: -100, max: -2, attempts: 50, err: nil},
		{min: -100, max: -102, attempts: 10, err: ErrInvalidRange},
	}
	for _, v := range tests {
		gen := NewIntGenerator(time.Now().UnixNano())
		allVals := make(map[int]bool)
		for i := 0; i < v.attempts; i++ {
			res, err := gen.Unique(v.min, v.max)
			if err != nil {
				if err != v.err {
					t.Fatalf("expected error: %v, got: %v", v.err, err)
				}
			} else {
				if _, exist := allVals[res]; exist {
					t.Fatalf("got duplicate value %d for range: %d - %d", res, v.min, v.max)
				}
				allVals[res] = true
			}
		}
	}
}
