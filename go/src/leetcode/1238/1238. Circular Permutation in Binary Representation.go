package main

/**
https://leetcode.cn/problems/circular-permutation-in-binary-representation/description/
思路：格雷编码的生成过程, G(i) = i ^ (i/2);
time: O(1<<n)
space: O(1<<n)
*/
func circularPermutation(n int, start int) []int {
	var ans = make([]int, 1<<n)
	for i := range ans {
		ans[i] = i ^ i>>1
		ans[i] = ans[i] ^ start
	}
	return ans
}

