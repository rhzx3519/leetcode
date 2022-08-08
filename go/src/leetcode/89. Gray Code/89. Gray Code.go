/**
@author ZhengHao Lou
Date    2022/01/08
*/
package main

/**
https://leetcode-cn.com/problems/gray-code/
思路：格雷编码的生成过程, G(i) = i ^ (i/2);
time: O(1<<n)
space: O(1<<n)
 */
func grayCode(n int) []int {
	ans := []int{}
	for i := 0; i < 1<<n; i++ {
		ans = append(ans, i^(i>>1))
	}
	return ans
}
