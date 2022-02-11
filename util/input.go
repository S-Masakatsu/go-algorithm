package util

import (
	"math/rand"
	"time"
)

type intRange struct {
	min, max int
}

func newIntRange(min, max int) *intRange {
	rand.Seed(time.Now().UnixNano())
	return &intRange{min, max}
}

func (ir *intRange) Intn() int {
	return rand.Intn(ir.max-ir.min+1) + ir.min
}

func RandSlice(size int) []int {
	r := newIntRange(-100, 100)
	ret := make([]int, size)
	for i := 0; i < len(ret); i++ {
		ret[i] = r.Intn()
	}
	return ret
}
