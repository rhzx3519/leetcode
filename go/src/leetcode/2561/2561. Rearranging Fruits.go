package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
https://leetcode.cn/problems/rearranging-fruits/
*/
func minCost(basket1 []int, basket2 []int) (tot int64) {
	n := len(basket1)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		cnt[basket1[i]]++
		cnt[basket2[i]]--
	}
	var mi = math.MaxInt32
	var a, b []int
	for x, v := range cnt {
		if v&1 == 1 {
			return -1
		}
		mi = min(mi, x)
		if v > 0 {
			for i := 0; i < v>>1; i++ { // a, b 需要预留一半给交换
				a = append(a, x)
			}
		} else if v < 0 {
			for i := 0; i < (-v)>>1; i++ {
				b = append(b, x)
			}
		}
	}

	sort.Ints(a)
	sort.Ints(b)
	reverse(b)
	for i := 0; i < len(a); i++ {
		tot += int64(min(a[i], min(b[i], mi<<1)))
	}

	fmt.Println(a, b, tot)
	return
}

func reverse(b []int) {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minCost([]int{4, 2, 2, 2}, []int{1, 4, 1, 2})
	minCost([]int{2, 3, 4, 1}, []int{3, 2, 5, 1})
}
