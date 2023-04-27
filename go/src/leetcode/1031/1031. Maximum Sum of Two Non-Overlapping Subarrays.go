package main

/*
*
https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/
*/
func maxSumTwoNoOverlap(nums []int, L1 int, L2 int) (tot int) {
	n := len(nums)
	sums := make([]int, n+1)
	for i := range nums {
		sums[i+1] = sums[i] + nums[i]
	}
	s := func(r, l int) int {
		return sums[min(n, r)] - sums[max(0, l)]
	}

	var l, lMax, r, rMax int
	for i := 0; i < n-L1-L2+1; i++ {
		l = max(l, s(i+L1, i))
		lMax = max(lMax, l+s(i+L1+L2, i+L1))
	}

	for i := 0; i < n-L1-L2+1; i++ {
		r = max(r, s(i+L2, i))
		rMax = max(rMax, r+s(i+L2+L1, i+L2))
	}
	return max(lMax, rMax)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxSumTwoNoOverlap([]int{0, 6, 5, 2, 2, 5, 1, 9, 4}, 1, 2)
}
