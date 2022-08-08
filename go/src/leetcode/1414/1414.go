/**
@author ZhengHao Lou
Date    2022/02/03
*/
package main

/**
https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
思路: 贪心

首先找到所有不超过 k 的斐波那契数字，然后每次贪心地选取不超过 k 的最大斐波那契数字，将 k 减去该斐波那契数字，
重复该操作直到 k 变为 0，此时选取的斐波那契数字满足和为 k 且数目最少。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/solution/he-wei-k-de-zui-shao-fei-bo-na-qi-shu-zi-shu-mu-by/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
time: O(f(k))
space: O(f(k))
*/
func findMinFibonacciNumbers(k int) int {
	if k <= 2 {
		return 1
	}
	var f1, f2 = 1, 1
	var f3 = f1 + f2
	nums := []int{f1, f2}
	for f3 <= k {
		f3 = f1 + f2
		nums = append(nums, f3)
		f1 = f2
		f2 = f3
	}

	var c int
	for i := len(nums) - 1; i >= 0 && k != 0; i-- {
		if k >= nums[i] {
			k -= nums[i]
			c++
		}
	}

	return c
}
