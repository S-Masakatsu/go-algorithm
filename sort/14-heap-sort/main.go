// ⭐️ Heap Sort (ヒープソート/ツリーソート)
// ====================================
// Avg  : O(n log n)
// Best : O(n log n)
// Worst: O(n log n)
// 安定 : No
package main

import (
	"fmt"
	"sort"

	"github/S-Masakatsu/go-algorithm/util"
)

func main() {
	input := util.RandSlice(10)
	fmt.Println("before:", input)
	heapSort(input)
	fmt.Println("after: ", input)
}

func heapSort(nums []int) {
	slice := sort.IntSlice(nums)
	for start := (slice.Len() - 2) / 2; start >= 0; start-- {
		siftDown(slice, start, slice.Len()-1)
	}

	for end := slice.Len() - 1; 0 < end; end-- {
		slice.Swap(0, end)
		siftDown(slice, 0, end-1)
	}
}

func siftDown(slice sort.Interface, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		if child+1 <= end && slice.Less(child, child+1) {
			child++
		}
		if !slice.Less(root, child) {
			return
		}
		slice.Swap(root, child)
		root = child
	}
}
