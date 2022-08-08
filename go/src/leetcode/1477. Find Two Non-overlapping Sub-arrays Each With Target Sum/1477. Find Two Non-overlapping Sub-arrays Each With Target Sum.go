/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/
f[i]表示i之前的最短子数组长度
 */
const INT_MAX int = math.MaxInt32 >> 1
func minSumOfLengths(arr []int, target int) int {
	n := len(arr)
	mapper := make(map[int]int)
	mapper[0] = 0

	var ans = INT_MAX
	f := make([]int, n+1)
	for i := range f {
		f[i] = INT_MAX
	}
	var s int
	for i := 1; i <= n; i++ {
		s += arr[i-1]
		f[i] = f[i - 1]
		if j, ok := mapper[s - target]; ok {
			ans = min(ans, f[j] + i - j)
			f[i] = min(f[i], i - j)
		}
		mapper[s] = i
	}
	fmt.Println(f, ans)
	if ans == INT_MAX {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minSumOfLengths([]int{3,2,2,4,3}, 3)
	minSumOfLengths([]int{7,3,4,7}, 7)
	minSumOfLengths([]int{3,1,1,1,5,1,2,1}, 3)
	minSumOfLengths([]int{5,5,4,4,5}, 3)
	minSumOfLengths([]int{4,3,2,6,2,3,4}, 6)
}
