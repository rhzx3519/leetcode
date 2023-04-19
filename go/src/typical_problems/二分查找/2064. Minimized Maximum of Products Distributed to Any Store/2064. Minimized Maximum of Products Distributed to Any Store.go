/**
@author ZhengHao Lou
Date    2021/11/08
*/
package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/minimized-maximum-of-products-distributed-to-any-store/
思路：二分查找
每间商店能够分配的商品数目介于0和max(quantities)之间
 */
func minimizedMaximum(n int, quantities []int) int {
	l, r := 1, max(quantities)
	var m int
	for l < r {
		m = l + (r-l)>>1	// 每间商店最多分配m件商品
		var c int			// 统计需要多少间商店
		for _, q := range quantities {
			c += (q + m - 1) / m		// if q <= m, c += 1; if q > m, c += >1
		}
		if c > n {
			l = m + 1
		} else {
			r = m
		}

	}
	return l
}

func max(a []int) int {
	var val int
	for _, num := range a {
		if num > val {
			val = num
		}
	}
	return val
}

func main() {
	fmt.Println(minimizedMaximum(6, []int{11,6}))
	fmt.Println(minimizedMaximum(7, []int{15,10,10}))
	fmt.Println(minimizedMaximum(1, []int{100000}))
}