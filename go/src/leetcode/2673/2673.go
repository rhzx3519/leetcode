package main

/*
*
https://leetcode.cn/problems/make-costs-of-paths-equal-in-a-binary-tree/description/
*/
func minIncrements(n int, cost []int) (tot int) {
	for i := n / 2; i > 0; i-- { // i是最后一个非叶节点
		left, right := cost[2*i-1], cost[2*i]
		if left > right {
			left, right = right, left
		}
		tot += right - left
		cost[i-1] += right
	}
	return
}
