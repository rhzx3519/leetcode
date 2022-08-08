/**
@author ZhengHao Lou
Date    2021/12/08
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
思路：区间和
 */
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	sums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	left := make([]int, n+1) // left[i] means max sum k in [:i]
	leftIndex := make([]int, n+1)
	right := make([]int, n+1) // right[i] means max sum k int[i:]
	rightIndex := make([]int, n+1)
	var s, maxS int
	for i := 0; i < n; i++ {
		if i >= k {
			if s > maxS {
				maxS = s
				leftIndex[i] = i - k
			} else {
				leftIndex[i] = leftIndex[i - 1]
			}
			left[i]	= maxS
			s -= nums[i-k]
		}
		s += nums[i]
	}
	s, maxS = 0, 0
	for i := n-1; i >= 0; i-- {
		s += nums[i]
		if i <= n-k {
			if s >= maxS {
				maxS = s
				rightIndex[i] = i
			} else {
				rightIndex[i] = rightIndex[i + 1]
			}
			right[i] = maxS
			s -= nums[i+k-1]
		}
	}

	fmt.Println(nums)
	fmt.Println(left)
	fmt.Println(leftIndex)
	fmt.Println(right)
	fmt.Println(rightIndex)

	s, maxS = 0, 0
	var ans = []int{}
	var l, r, m int
	for i := k; i + k<<1 <= n; i++ {		// (0,k-1) (k, k+k-1) (k+k, k+k+k-1)
		m = sums[i+k] - sums[i]
		l, r = left[i], right[i+k]
		s = l + m + r
		fmt.Println(l, m, r)
		if s > maxS {
			maxS = s
			ans = []int{leftIndex[i], i, rightIndex[i+k]}
		}
	}

	fmt.Println("ans:",  ans)

	return ans
}

func main() {
	maxSumOfThreeSubarrays([]int{1,2,1,2,6,7,5,1}, 2)		// [0,3,5]
	//maxSumOfThreeSubarrays([]int{1,2,1,2,1,2}, 2)	// [0,2,4]
	//maxSumOfThreeSubarrays([]int{4,3,2,1}, 1)	// [0,1,2]
}