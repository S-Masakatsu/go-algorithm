package util

import (
	"math/rand"
	"time"
)

func RandSlice(size int) []int {
	rand.Seed(time.Now().UnixNano())
	ret := make([]int, size)
	for i := 0; i < len(ret); i++ {
		ret[i] = rand.Intn(100)
	}
	return ret
}
