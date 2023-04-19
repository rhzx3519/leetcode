package main

/*
* https://leetcode.cn/problems/maximize-win-from-two-segments/solutions/2093246/tong-xiang-shuang-zhi-zhen-ji-lu-di-yi-t-5hlh/
思路：不固定窗口大小的，同向双指针算法
*/
func maximizeWin(prizePositions []int, k int) int {
	n := len(prizePositions)
	var pre = make([]int, n+1)
	var l, ans int
	for r, p := range prizePositions {
		for p-prizePositions[l] > k {
			l++
		}
		ans = max(ans, r-l+1+pre[l])
		pre[r+1] = max(pre[r], r-l+1)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
