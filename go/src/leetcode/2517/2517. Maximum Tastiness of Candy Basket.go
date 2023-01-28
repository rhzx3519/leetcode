/*
*
@author ZhengHao Lou
Date    2022/12/26
*/
package main

import "sort"

/*
*
https://leetcode.cn/problems/maximum-tastiness-of-candy-basket/
*/
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	return sort.Search((price[len(price)-1]+price[0])/(k-1), func(d int) bool {
		d++
		var c, x0 = 1, price[0]
		for _, x1 := range price[1:] {
			if x0+d <= x1 {
				c++
				x0 = x1
			}
		}
		return c < k
	})
}
