/**
@author ZhengHao Lou
Date    2022/03/31
*/
package main

/**
https://leetcode-cn.com/problems/self-dividing-numbers/
*/
func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		if is(i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func is(x int) bool {
	for i := x; i != 0; i /= 10 {
		d := i % 10
		if d == 0 || x%d != 0 {
			return false
		}
	}
	return true
}
