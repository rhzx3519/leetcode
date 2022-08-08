/**
@author ZhengHao Lou
Date    2021/11/07
*/
package main

/**
https://leetcode-cn.com/problems/range-addition-ii/
 */
func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		a, b := op[0], op[1]
		m = min(m, a)
		n = min(n, b)
	}
	return m * n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
