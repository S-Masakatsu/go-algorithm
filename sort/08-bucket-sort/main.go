// Bucket Sort (バケットソート)
// バケットという単位で配列を分割して挿入ソートをおこなう
// ====================================
// Avg  : O(n+k) k...リストのサイズで変化
// Best : O(n+k)
// Worst: O(n^2)
// 安定 : Yes
// 備考 : InsertionSortを使用
package main

import (
	"fmt"
	"math"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	ans := bucketSort(input)
	fmt.Println("after: ", ans)
}

func bucketSort(nums []int) []int {
	maxNum := max(nums...)
	size := maxNum / len(nums)
	bucket := make([][]int, size)

	for _, n := range nums {
		i := n / size
		if i < size {
			bucket[i] = append(bucket[i], n)
		} else {
			bucket[size-1] = append(bucket[size-1], n)
		}
	}

	ret := []int{}
	for _, b := range bucket {
		insertionSort(b)
		ret = append(ret, b...)
	}

	return ret
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

func max(nums ...int) int {
	ret := nums[0]
	for _, n := range nums {
		ret = int(math.Max(float64(ret), float64(n)))
	}
	return ret
}
