// Linear Search (線形探索)
// ====================================
// 最も基本的な検索アルゴリズム
// 基本的に、目的の結果が見つかり、そのインデックスを返すまで、配列を反復処理する
// デメリット：探したい値が後ろにあるほど遅い
package main

import "fmt"

func main() {
	list := []int{9, 1, 33, 21, 78, 42, 4}
	fmt.Println("list: ", list)
	fmt.Println("21 to index:", linearSearch(list, 21))
	fmt.Println("99 to index:", linearSearch(list, 99))
}

func linearSearch(nums []int, s int) int {
	for i, n := range nums {
		if n == s {
			return i
		}
	}
	return -1
}
