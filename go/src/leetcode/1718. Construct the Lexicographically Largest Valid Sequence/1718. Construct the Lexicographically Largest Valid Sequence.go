package main

import "fmt"

/**
https://leetcode-cn.com/problems/construct-the-lexicographically-largest-valid-sequence/
思路：贪心 + 回溯查找
贪心: 按照数字从大到小填充数组，第一个符合要求的数组必然是最大的。
回溯查找：使用回溯查找填充数组。
 */
func constructDistancedSequence(n int) []int {
	N := 2*n - 1
	arr := make([]int, N)

	vis := make([]bool, n + 1)
	var ans []int

	var dfs func(int, []int, int) bool
	dfs = func(idx int, arr[]int, filled int) bool {
		if filled == N {
			ans = append(ans, arr...)
			return true
		}

		if arr[idx] != 0 {
			return dfs(idx + 1, arr, filled)
		}

		for i := n; i >= 1; i-- {
			if vis[i] {
				continue
			}
			if i > 1 && (idx + i >= N || arr[idx + i] != 0)	{
				continue
			}

			vis[i] = true
			arr[idx] = i
			filled++
			if i > 1 {
				arr[idx + i] = i
				filled++
			}

			if dfs(idx + 1, arr, filled) {
				return true
			}

			arr[idx] = 0
			filled--
			if i > 1 {
				arr[idx + i] = 0
				filled--
			}
			vis[i] = false
		}
		return false
	}

	dfs(0, arr, 0)
	return ans
}

func main() {
	fmt.Println(constructDistancedSequence(1))
	fmt.Println(constructDistancedSequence(2))
	fmt.Println(constructDistancedSequence(3))
	fmt.Println(constructDistancedSequence(5))
}