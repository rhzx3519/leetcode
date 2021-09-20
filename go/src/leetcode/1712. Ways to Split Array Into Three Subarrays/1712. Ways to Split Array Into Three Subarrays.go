package main

import "fmt"

/**
题目：https://leetcode-cn.com/problems/ways-to-split-array-into-three-subarrays/
题解：https://leetcode-cn.com/problems/ways-to-split-array-into-three-subarrays/solution/5643-jiang-shu-zu-fen-cheng-san-ge-zi-sh-fmep/
思路：i, j 表示分割点, [0:i), [i:j), [j:n)
使用二分寻找满足要求的最左边的j1，和最右边的j2,
对于区间[i:], j越往右，则mid = sum[i:j)越大，right = sum[j:n)越小
 */

const MOD int = 1e9 + 7

func waysToSplit(nums []int) int {
	n := len(nums)
	sum := make([]int, n)
	sum[0] = nums[0]
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	s3 := sum[n-1] / 3
	var count int
	var left, right int
	for i := 0; i < n; i++ {
		if sum[i] > s3 {
			break
		}
		left = lowerBound(sum, i+1, n, sum[i] << 1)		// 最左边的分割点，满足mid >= left
		if left == n {
			// 不存在分割点满足要求
			continue
		}
		right = upperBound(sum, i+1, n, sum[i] + (sum[n-1] - sum[i]) / 2)	// 最右边的分割点j， 满足right >= mid
		if right == n {
			// right子数组需要非空
			right--
		}
		if left <= right {
			count = (count + right - left) % MOD
		}
	}

	return count
}

func lowerBound(nums []int, l, r, x int) int {
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func upperBound(nums []int, l, r, x int) int {
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}



func main() {
	fmt.Println(waysToSplit([]int{1,1,1}))
	fmt.Println(waysToSplit([]int{1,2,2,2,5,0}))
	fmt.Println(waysToSplit([]int{3,2,1}))
}

