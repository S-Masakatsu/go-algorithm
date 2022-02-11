package util

import (
	"math/rand"
	"sort"
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

func RandSearchSlice(size int, isSort bool) ([]int, int) {
	nums := RandSlice(size)
	if isSort {
		sort.Ints(nums)
	}

	min, max := Min(nums...), Max(nums...)
	search := rand.Intn(max-min+1) + min
	// ↑だけだと存在しない値ばかりになるので共済処置
	if i := rand.Intn(size + 10); i < len(nums) {
		search = nums[i]
	}

	return nums, search
}
