// ⭐️ Selection Sort (選択ソート)
// 最小値を保持し、先頭にソートする
// ====================================
// Avg  : O(n^2)
// Best : O(n^2)
// Worst: O(n^2)
// 安定 : Yes
// 備考 : BubbleSortの改良
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	selectionSort(input)
	fmt.Println("after: ", input)
}

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		tmpMin := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[tmpMin] {
				tmpMin = j
			}
		}
		nums[i], nums[tmpMin] = nums[tmpMin], nums[i]
	}
}
