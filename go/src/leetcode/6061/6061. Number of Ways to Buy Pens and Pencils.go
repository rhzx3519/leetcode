/**
@author ZhengHao Lou
Date    2022/04/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/number-of-ways-to-buy-pens-and-pencils/
*/
func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	if total < cost1 && total < cost2 {
		return 1
	}
	var c int64
	for i := 0; cost1*i <= total; i++ {
		left := total - cost1*i
		c2 := left/cost2 + 1

		c += int64(c2)
	}

	fmt.Println(c)
	return c
}

func main() {
	waysToBuyPensPencils(20, 10, 5)
	waysToBuyPensPencils(5, 10, 10)
}
