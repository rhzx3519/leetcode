/**
@author ZhengHao Lou
Date    2022/01/15
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/calculate-money-in-leetcode-bank/
 */
func totalMoney(n int) int {
	var money int
	var i = 1
	for ; i <= n/7; i++ {
		money += sum(i, 7)
	}
	money += sum(i, n%7)
	return money
}

func sum(a1, n int) int {
	return n*(a1 + a1 + n - 1)/2
}

func main() {
	fmt.Println(totalMoney(20))
}