package main

/*
* https://leetcode.cn/problems/the-number-of-beautiful-subsets/
思路：回溯查找
*/
func beautifulSubsets(nums []int, k int) int {
	tot := -1
	n := len(nums)
	cnt := make(map[int]int)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			tot++
			return
		}
		dfs(i + 1)
		x := nums[i]
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++
			dfs(i + 1)
			cnt[x]--
		}
	}

	return tot
}
