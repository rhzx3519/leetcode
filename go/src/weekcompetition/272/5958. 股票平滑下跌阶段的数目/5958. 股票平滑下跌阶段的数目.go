/**
@author ZhengHao Lou
Date    2021/12/19
*/
package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	var c int64
	//n := len(prices)
	var s int
	for i := range prices {
		if i > 0 && prices[i] == prices[i-1] - 1 {
			s++
		} else {
			s = 1
		}
		c += int64(s)
		fmt.Println(s)
	}

	fmt.Println("c:", c)
	return c
}

func main() {
	//getDescentPeriods([]int{3,2,1,4})
	//getDescentPeriods([]int{8,6,7,7})
	getDescentPeriods([]int{1})
}