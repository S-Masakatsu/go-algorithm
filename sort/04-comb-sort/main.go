// Comb Sort (コムソート)
// コム(串) 串を溶かすイメージで串の幅を小さくしていきながらソートする
// Gap(串)が1 && swapがfalseならソート完了
// ====================================
// Avg  : O(n^2/2**g) (g is a number of increment) Gapの値で変化する
// Best : O(n log n)
// Worst: O(n^2)
// 安定 : No
// 備考 : BubbleSortの改良(場合によっては劣化している)
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	combSort(input)
	fmt.Println("after: ", input)
}

func combSort(nums []int) {
	for gap := len(nums); ; {
		if 1 < gap {
			gap = gap * 4 / 5 // gap = gap / 1.3
		}
		swapped := false
		for i := 0; i < len(nums)-gap; i++ {
			if nums[i] > nums[i+gap] {
				nums[i], nums[i+gap] = nums[i+gap], nums[i]
				swapped = true
			}
		}

		if gap == 1 && !swapped {
			return
		}
	}
}
