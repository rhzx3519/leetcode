package main

import (
	"fmt"
	"sort"
)

/*
* https://leetcode.cn/problems/maximum-fruits-harvested-after-at-most-k-steps/description/
思路：贪心，二分搜索
最优的搜索方式是往左/往右走一段距离，然后再返回往右/往左走一段距离，正好加起来等于k步，这样可以在k步内获取最多的水果，
枚举左右分别走的步数
*/
func maxTotalFruits(fruits [][]int, startPos int, k int) (ans int) {
	n := len(fruits)
	sums := make([]int, n+1)
	pos := make([]int, n)
	for i, x := range fruits {
		sums[i+1] = sums[i] + x[1]
		pos[i] = x[0]
	}

	for x := k; x >= 0; x-- {
		y := k - x<<1
		l := sort.SearchInts(pos, startPos-x)
		r := sort.SearchInts(pos, startPos+y+1)
		ans = max(ans, sums[r]-sums[l])

		l = sort.SearchInts(pos, startPos-y)
		r = sort.SearchInts(pos, startPos+x+1)
		ans = max(ans, sums[r]-sums[l])
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxTotalFruits([][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4))
}
