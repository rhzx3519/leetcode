/**
@author ZhengHao Lou
Date    2022/02/21
*/
package main

/**
https://leetcode-cn.com/problems/count-integers-with-even-digit-sum/
*/
func countEven(num int) int {
	var c int
	for i := 1; i <= num; i++ {
		var s int
		for x := i; x != 0; x /= 10 {
			s += x % 10
		}
		if s&1 == 0 {
			c++
		}
	}
	return c
}
