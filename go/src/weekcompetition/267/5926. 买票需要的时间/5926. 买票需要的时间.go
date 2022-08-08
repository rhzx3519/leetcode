/**
@author ZhengHao Lou
Date    2021/11/14
*/
package main

import "fmt"

func timeRequiredToBuy(tickets []int, k int) int {
	var t int
	n := len(tickets)
	var i int
	for tickets[k] != 0 {
		if tickets[i] != 0 {
			tickets[i]--
			t++
		}
		i = (i + 1)%n
	}
	return t
}

func main() {
	fmt.Println(timeRequiredToBuy([]int{2,3,2}, 2))
}
