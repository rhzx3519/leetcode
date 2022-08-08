/**
@author ZhengHao Lou
Date    2022/01/09
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-swaps-to-group-all-1s-together-ii/
思路：滑动窗口，前缀和
n为数组中元素个数，k为1的个数，最少的交换次数为长度为k的子数组中0的个数的最小值
time: O(n)
space: O(n)
 */
func minSwaps(nums []int) int {
	var k int
	for i := range nums {
		if nums[i] == 1 {
			k++
		}
	}

	nums = append(nums, nums...)
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}

	var ans = math.MaxInt32 >> 1
	var c int
	for i := k; i <= n; i++ {
		c = sums[i] - sums[i-k]
		ans = min(ans, k - c)
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
	fmt.Println(minSwaps([]int{0,1,0,1,1,0,0}))
	fmt.Println(minSwaps([]int{0,1,1,1,0,0,1,1,0}))
	fmt.Println(minSwaps([]int{1,1,0,0,1}))
}
