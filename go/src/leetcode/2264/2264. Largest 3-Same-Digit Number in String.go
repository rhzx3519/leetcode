/**
@author ZhengHao Lou
Date    2022/05/09
*/
package main

/**
https://leetcode.cn/problems/largest-3-same-digit-number-in-string/
*/
func largestGoodInteger(num string) string {
	var x string
	var c int
	var b byte
	for i := range num {
		if num[i] == b {
			c++
		} else {
			c = 1
		}
		if c >= 3 && num[i-2:i+1] > x {
			x = num[i-2 : i+1]
		}
		b = num[i]

	}
	return x
}
