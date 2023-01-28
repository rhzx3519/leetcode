/*
*
@author ZhengHao Lou
Date    2022/12/26
*/
package main

/*
*
https://leetcode.cn/problems/number-of-great-partitions/description/
思路：逆向思维 + 01背包
f[j]表示分区和等于j的分区数量
*/
const mod int = 1e9 + 7

func countPartitions(nums []int, k int) int {
	var s int
	for i := range nums {
		s += nums[i]
	}
	if s < k<<1 {
		return 0
	}

	var ans = 1
	f := make([]int, k)
	f[0] = 1
	for _, x := range nums {
		ans = ans * 2 % mod
		for j := k - 1; j >= x; j-- {
			f[j] = (f[j] + f[j-x]) % mod
		}
	}
	for i := range f {
		ans -= 2 * f[i]
	}
	return (ans%mod + mod) % mod
}
