/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts/description/
https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts/solutions/699661/cong-zui-zhi-jie-de-fang-fa-kai-shi-yi-b-x729/

*/
func maxHappyGroups(batchSize int, groups []int) int {
	cnt := make([]int, batchSize)
	for i := range groups {
		cnt[groups[i]%batchSize]++
	}
	var all = 1
	for i := 1; i < batchSize; i++ {
		all *= cnt[i] + 1
	}
	f := make([]int, all)
	for state := 1; state < all; state++ {
		var sum = batchSize
		for n, t := 1, state; n < batchSize; t, n = t/(cnt[n]+1), n+1 {
			sum += n * (t % (cnt[n] + 1))
		}
		for n, t, w := 1, state, 1; n < batchSize; t, n, w = t/(cnt[n]+1), n+1, w*(cnt[n]+1) {
			if t%(cnt[n]+1) != 0 {
				f[state] = max(f[state], f[state-w]+b2i((sum-n)%batchSize == 0))
			}
		}
	}
	return f[all-1] + cnt[0]
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxHappyGroups(4, []int{1, 3, 2, 5, 2, 2, 1, 6}))
}
