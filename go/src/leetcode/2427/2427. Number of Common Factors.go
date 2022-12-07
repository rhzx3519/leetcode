/**
@author ZhengHao Lou
Date    2022/10/05
*/
package main

/**
https://leetcode.cn/problems/number-of-common-factors/
*/
func commonFactors(a int, b int) int {
	var c int
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			c++
		}
	}
	return c
}
