package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/largest-rectangle-in-histogram/
类型leetcode-84, 使用前缀和和单调栈。
 */
var MOD = int64(math.Pow10(9)) + 7

func maxSumMinProduct(nums []int) int {
	a := []int{0}
	a = append(a, nums...)
	a = append(a, 0)

	n := len(a)

	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + a[i-1]
	}

	var ans int64
	st := []int{}
	for i := 0; i < n; i++ {
		for len(st) > 0 && a[st[len(st)-1]] > a[i] {
			last := st[len(st)-1]
			st = st[:len(st)-1]
			ans = max(ans, int64(a[last]*(sum[i] - sum[st[len(st)-1]+1])))
		}
		st = append(st, i)
	}
	return int(ans % MOD)
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSumMinProduct([]int{1,2,3,2}))
	fmt.Println(maxSumMinProduct([]int{2,3,3,1,2}))
	fmt.Println(maxSumMinProduct([]int{3,1,5,6,4,2}))
}