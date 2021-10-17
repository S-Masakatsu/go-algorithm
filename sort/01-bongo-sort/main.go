// Bongo Sort (ボゴソート)
// Avg  : O((n+1)!)
// Best : O(n)
// Worst: Unbounded
// 安定 : No
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	bongoSort(input)
	fmt.Println("after: ", input)
}

func bongoSort(nums []int) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	if !isOrder(nums) {
		bongoSort(nums)
	}
}

func isOrder(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}
