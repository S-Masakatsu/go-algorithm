// Cocktail Sort (カクテルソート/シェーカーソート)
// 左右にシャッフルしながら隣同士比較して交換する
// ====================================
// Avg  : O(n^2)
// Best : O(n)
// Worst: O(n^2)
// 安定 : Yes
// 備考 : BubbleSortの改良(多少早い)
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	cocktailSort(input)
	fmt.Println("after: ", input)
}

func cocktailSort(nums []int) {
	for start, end := 0, len(nums)-1; ; {
		swapped := false
		for i := start; i < end; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}

		if !swapped {
			return
		}
		end--
		swapped = false

		for i := end - 1; i >= start; i-- {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}

		if !swapped {
			return
		}
		start++
	}
}
