/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

func alternateDigitSum(n int) int {
	var a []int
	for ; n != 0; n /= 10 {
		a = append(a, n%10)
	}
	var ans int
	var k = 1
	for i := len(a) - 1; i >= 0; i-- {
		ans += a[i] * k
		k = -k
	}
	return ans
}
