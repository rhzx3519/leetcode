/**
@author ZhengHao Lou
Date    2022/06/12
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/
*/
func countSubarrays(nums []int, k int64) (ans int64) {
	var l int
	var s int64
	for r, num := range nums {
		s += int64(num)
		for l < r && s*int64(r-l+1) >= k {
			s -= int64(nums[l])
			l++
		}
		ans += int64(r - l + 1)
	}
	fmt.Println(ans)
	return ans
}

func main() {
	countSubarrays([]int{2, 1, 4, 3, 5}, 10)
}
