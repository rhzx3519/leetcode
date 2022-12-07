/*
*
@author ZhengHao Lou
Date    2022/11/27
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/minimum-penalty-for-a-shop/
*/
func bestClosingTime(customers string) int {
	var tot int
	for i := range customers {
		switch customers[i] {
		case 'Y':
			tot++
		case 'N':
		
		}
	}
	var j, cost = 0, tot
	var c int
	for i := range customers {
		switch customers[i] {
		case 'Y':
			tot--
		case 'N':
			c++
		}
		if tot+c < cost {
			cost = tot + c
			j = i + 1
		}
	}
	
	fmt.Println(cost, j)
	return j
}

func main() {
	bestClosingTime("YYNY")
	bestClosingTime("NNNNN")
	bestClosingTime("YYYY")
}
