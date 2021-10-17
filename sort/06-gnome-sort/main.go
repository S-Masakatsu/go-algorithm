// Gnome Sort (ノームソート)
// バブルソートと同じだがソート(入れ替えたら)１つ戻る
// 最後の要素までいったらソート完了
// ====================================
// Avg  : O(n^2)
// Best : O(n)
// Worst: O(n^2)
// 安定 : Yes
// 備考 : BubbleSortに類似
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	gnomeSort(input)
	fmt.Println("after: ", input)
}

func gnomeSort(nums []int) {
	for i, j := 1, 2; i < len(nums); {
		if nums[i-1] > nums[i] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i--
			if i > 0 {
				continue
			}
		}
		i = j
		j++
	}
}
