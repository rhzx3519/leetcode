/**
@author ZhengHao Lou
Date    2022/10/26
*/
package main

/**
https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/
https://leetcode.cn/problems/shortest-subarray-with-sum-at-least-k/solution/liang-zhang-tu-miao-dong-dan-diao-dui-li-9fvh/#comment
*/
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	var ans = n + 1
	var que []int
	for i, cs := range sums {
		for len(que) > 0 && cs-sums[que[0]] >= k {
			ans = min(ans, i-que[0])
			que = que[1:]
		}
		for len(que) > 0 && sums[que[len(que)-1]] >= cs {
			que = que[:len(que)-1]
		}
		que = append(que, i)
	}
	if ans > n {
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
