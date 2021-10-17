// ⭐️ Insertion Sort (挿入ソート)
// 今見ている値を適切な位置に持っていきながらソートする
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
	insertionSort(input)
	fmt.Println("after: ", input)
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		j, val := i-1, nums[i]
		for ; j >= 0 && nums[j] > val; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = val
	}
}
