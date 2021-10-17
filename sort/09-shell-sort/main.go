// Shell Sort (シェルソート)
// ====================================
// Avg  : Depends on gap sequence
// Best : O(n log n)
// Worst: O(n^2)
// 安定 : No
// 備考 : InsertionSortの改良
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	shellSort(input)
	fmt.Println("after: ", input)
}

func shellSort(nums []int) {
	for gap := len(nums) / 2; 0 < gap; gap /= 2 {
		for i := gap; i < len(nums); i++ {
			j, val := i, nums[i]
			for ; j >= gap && nums[j-gap] > val; j -= gap {
				nums[j] = nums[j-gap]
			}
			nums[j] = val
		}
	}
}
