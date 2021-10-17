// Radix Sort (基数ソート)
// 数を分割してCounttingSortをする
// 1の位, 10の位 ... Nの位をCountingSortする
// ====================================
// Avg  : O(n)
// Best : O(n)
// Worst: O(n)
// 安定 : Yes
// 備考 : CountingSortを使用
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := make([]int32, 10)
	for i, n := range util.RandSlice(10) {
		input[i] = int32(n)
	}
	fmt.Println("before:", input)
	radixSort(input)
	fmt.Println("after: ", input)
}

func radixSort(nums []int32) {
	const wordLen = 4
	const bitHigh = -1 << 31

	buf := &bytes.Buffer{}
	ds := make([][]byte, len(nums))
	for i, n := range nums {
		binary.Write(buf, binary.LittleEndian, n^bitHigh)
		b := make([]byte, wordLen)
		buf.Read(b)
		ds[i] = b
	}

	bins := make([][][]byte, 256)
	for i := 0; i < wordLen; i++ {
		for _, b := range ds {
			bins[b[i]] = append(bins[b[i]], b)
		}

		j := 0
		for key, bs := range bins {
			copy(ds[j:], bs)
			j += len(bs)
			bins[key] = bs[:0]
		}
	}

	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		nums[i] = w ^ bitHigh
	}
}
