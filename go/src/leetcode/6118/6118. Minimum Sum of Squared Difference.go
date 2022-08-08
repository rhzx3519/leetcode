/**
@author ZhengHao Lou
Date    2022/07/10
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/minimum-sum-of-squared-difference/
思路：贪心，
每次把最大的差值-1
*/
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	n := len(nums1)
	var counter = make(map[int]int)
	var nums = make([]int, n)
	var mx int
	for i := range nums1 {
		nums[i] = abs(nums2[i] - nums1[i])
		if nums[i] > mx {
			mx = nums[i]
		}
		counter[nums[i]]++
	}

	for k := k1 + k2; k != 0; {
		if mx == 0 {
			return 0
		}
		if k > counter[mx] {
			counter[mx-1] += counter[mx]
			k -= counter[mx]
			delete(counter, mx)
			mx--
		} else {
			counter[mx] -= k
			counter[mx-1] += k
			k = 0
		}
	}

	var ans int64
	for i := range counter {
		ans += int64(i) * int64(i) * int64(counter[i])
	}

	fmt.Println(ans)
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	minSumSquareDiff([]int{1, 2, 3, 4}, []int{2, 10, 20, 19}, 0, 0)
	minSumSquareDiff([]int{1, 4, 10, 12}, []int{5, 8, 6, 9}, 1, 1)
	// 27
	minSumSquareDiff([]int{7, 11, 4, 19, 11, 5, 6, 1, 8}, []int{4, 7, 6, 16, 12, 9, 10, 2, 10}, 3, 6)
}
