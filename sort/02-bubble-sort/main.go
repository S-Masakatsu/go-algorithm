// ⭐️ Bubble Sort (バブルソート)
// 隣同士比較して交換するだけ
// ====================================
// Avg  : O(n^2)
// Best : O(n)
// Worst: O(n^2)
// 安定 : Yes
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	bubbleSort(input)
	fmt.Println("after: ", input)
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
