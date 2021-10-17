// Counting Sort (計数ソート)
// 1) 最大値の要素分の配列を用意
// 2) 値のカウントを取る
// 3) 自分の前の要素に何個数字があるか確かめる
// ====================================
// Avg  : O(n)
// Best : O(n)
// Worst: O(n)
// 安定 : Yes
// 備考 : 最大値がものすごい大きい場合はメモリを食う
package main

import (
	"fmt"
	"math"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	countingSort(input)
	fmt.Println("after: ", input)
}

func countingSort(nums []int) {
	minNum, maxNum := min(nums...), max(nums...)
	size := maxNum - minNum + 1
	count := make([]int, size)

	for _, n := range nums {
		count[n-minNum]++
	}

	z := 0
	for i, cnt := range count {
		for ; cnt > 0; cnt-- {
			nums[z] = i + minNum
			z++
		}
	}
}

func min(nums ...int) int {
	ret := nums[0]
	for _, n := range nums {
		ret = int(math.Min(float64(ret), float64(n)))
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for _, n := range nums {
		ret = int(math.Max(float64(ret), float64(n)))
	}
	return ret
}
