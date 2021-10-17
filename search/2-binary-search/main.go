// Binary Search (二分探索)
// ====================================
// 有名で高速な検索アルゴリズム
// 昇順でソートされた配列を提供する必要がある
// この並べ替えでは、基本的に配列の中間値を調べて、それがターゲットと等しいかどうかを確認する。
// 中央の値がターゲットよりも小さい場合は、ターゲットが中央の値を超えていることを意味し、
// 大きい場合は、ターゲットが中央の値の前にあることを意味する。
// また、アルゴリズムは、ターゲットと一致するまで中間値を見つける傾向がある。
// LinearSearchのデメリットを解消できる
package main

import "fmt"

func main() {
	list := []int{2, 9, 11, 21, 22, 32, 36, 48, 76}
	fmt.Println("list: ", list)
	fmt.Println("21 to index:", binarySearch(list, 21))
	fmt.Println("99 to index:", binarySearch(list, 99))
	fmt.Println("32 to index:", recursiveBinarySearch(list, 32))
}

func binarySearch(nums []int, s int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == s:
			return mid
		case nums[mid] < s:
			left = mid + 1
		default:
			right = mid - 1
		}
	}
	return -1
}

// 再帰を使った書き方
func recursiveBinarySearch(nums []int, s int) int {
	var search func(nums []int, s, left, right int) int
	search = func(nums []int, s, left, right int) int {
		if left > right {
			return -1
		}

		mid := (left + right) / 2
		switch {
		case nums[mid] == s:
			return mid
		case nums[mid] < s:
			return search(nums, s, mid+1, right)
		default:
			return search(nums, s, left, mid-1)
		}
	}

	return search(nums, s, 0, len(nums)-1)
}
