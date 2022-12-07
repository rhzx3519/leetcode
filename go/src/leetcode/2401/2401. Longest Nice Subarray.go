/**
@author ZhengHao Lou
Date    2022/09/05
*/
package main

/**
https://leetcode.cn/problems/longest-nice-subarray/
*/
func longestNiceSubarray(nums []int) int {
	n := len(nums)
	var ans int
	var l, r, mask int
	for ; r < n; r++ {
		for l < r && mask&nums[r] != 0 {
			mask ^= nums[l]
			l++
		}
		mask ^= nums[r]
		ans = max(ans, r-l+1)
	}
	//fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	longestNiceSubarray([]int{1, 3, 8, 48, 10})
	longestNiceSubarray([]int{3, 1, 5, 11, 13})
}
