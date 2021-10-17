// Jump Search (ジャンプ探索)
// ====================================
// ソートされた配列に使用されるアルゴリズム
// ターゲット以上の値が見つかるまで、配列の長さの平方根でジャンプする。
// 逆線形検索を実装して、最終結果を決定する必要がある
package main

import (
	"fmt"
	"math"
)

func main() {
	list := []int{2, 9, 11, 21, 22, 32, 36, 48, 76, 90}
	fmt.Println("list: ", list)
	fmt.Println("21 to index:", jumpSearch(list, 21))
	fmt.Println("99 to index:", jumpSearch(list, 99))
}

func jumpSearch(nums []int, s int) (ret int) {
	ret = -1
	defer func() {
		if recover() != nil {
			return
		}
	}()

	size := int(math.Sqrt(float64(len(nums))))
	i := 0
	for ; ; i += size {
		if nums[i] >= s {
			break
		} else if i > len(nums) {
			break
		}
	}

	for j := i; j > 0; j-- {
		if nums[j] == s {
			return j
		}
	}
	return
}
