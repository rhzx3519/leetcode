/**
@author ZhengHao Lou
Date    2022/05/18
*/
package main

/**
https://leetcode.cn/problems/kth-smallest-number-in-multiplication-table/
*/
func findKthNumber(m int, n int, k int) int {
	l, r := 1, m*n+1
	for l < r {
		mid := l + (r-l)>>1
		c := mid / n * n
		for i := mid/n + 1; i <= m; i++ {
			c += mid / i
		}
		if c >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	//findKthNumber(3, 3, 5)
	findKthNumber(2, 3, 6)
}
