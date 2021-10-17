// ⭐️ Quick Sort (クイックソート)
// ====================================
// Avg  : O(n log n)
// Best : O(n log n)
// Worst: O(n^2)
// 安定 : No
package main

import (
	"fmt"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	quickSort(input)
	fmt.Println("after: ", input)
}

func quickSort(nums []int) {
	var sort func(nums []int, low, high int)
	sort = func(nums []int, low, high int) {
		if low < high {
			q := partition(nums, low, high)
			sort(nums, low, q-1)
			sort(nums, q+1, high)
		}
	}

	sort(nums, 0, len(nums)-1)
}

func partition(nums []int, low, high int) int {
	i := low - 1
	pivot := nums[high]
	for j := low; j < high; j++ {
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	i++
	nums[i], nums[high] = nums[high], nums[i]
	return i
}
