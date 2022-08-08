/**
@author ZhengHao Lou
Date    2021/11/22
*/
package main

/**
https://leetcode-cn.com/problems/factorial-trailing-zeroes/
思路：求n的因式分解中5的个数
 */
func trailingZeroes(n int) int {
	var count int
	for n >= 5 {
		count += n/5
		n /= 5
	}
	return count
}
