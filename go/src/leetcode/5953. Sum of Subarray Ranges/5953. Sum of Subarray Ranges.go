/**
@author ZhengHao Lou
Date    2021/12/13
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/sum-of-subarray-ranges/
思路：数据量少，暴力
time: O(n^2)
space: O(1)
 */
func subArrayRanges(nums []int) int64 {
	var ans int64
	n := len(nums)
	for i := 0; i < n; i++ {
		var mx, mi int64 = int64(nums[i]), int64(nums[i])
		for j := i; j < n; j++ {
			if int64(nums[j]) > mx {
				mx = int64(nums[j])
			}
			if int64(nums[j]) < mi {
				mi = int64(nums[j])
			}
			ans += mx - mi
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	subArrayRanges([]int{1,2,3})
	subArrayRanges([]int{1,3,3})
	subArrayRanges([]int{4,-2,-3,4,1})
}