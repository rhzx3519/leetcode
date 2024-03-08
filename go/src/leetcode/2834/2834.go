package main

/*
*
https://leetcode.cn/problems/find-the-minimum-possible-sum-of-a-beautiful-array/description/?envType=daily-question&envId=2024-03-08
*/
func minimumPossibleSum(n int, k int) int {
	m := min(k/2, n)
	return (m*(m+1) + ((2*k + n - m - 1) * (n - m))) / 2 % int(1e9+7)
}
