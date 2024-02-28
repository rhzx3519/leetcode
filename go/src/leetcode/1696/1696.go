package main

/*
*
https://leetcode.cn/problems/jump-game-vi/?envType=daily-question&envId=2024-02-05
*/
func maxResult(nums []int, k int) int {
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	q := []int{0}
	for i := 1; i < n; i++ {
		if i-k > q[0] {
			q = q[1:]
		}
		f[i] = nums[i] + f[q[0]]
		for len(q) != 0 && f[i] >= f[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	return f[n-1]
}
