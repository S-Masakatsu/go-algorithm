// ⭐️ Merge Sort (マージソート)
// ====================================
// 最後の１つになるまで分割して、マージする時にソートする
// Avg  : O(n log n)
// Best : O(n log n)
// Worst: O(n log n)
// 安定 : Yes
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	ans := mergeSort(input)
	fmt.Println("after: ", ans)
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	var i, j int
	ret := []int{}
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			ret = append(ret, left[i])
			i++
		} else {
			ret = append(ret, right[j])
			j++
		}
	}

	ret = append(ret, left[i:]...)
	ret = append(ret, right[j:]...)

	return ret
}
