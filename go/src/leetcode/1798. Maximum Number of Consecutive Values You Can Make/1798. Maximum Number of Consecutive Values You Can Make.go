/**
https://leetcode-cn.com/problems/maximum-number-of-consecutive-values-you-can-make/

 */
package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode-cn.com/problems/maximum-number-of-consecutive-values-you-can-make/
思路：假设能够早0~4, 则下一个coin如果 <= 5, 则至少可以构造 4+coin的连续数字, 大于5则无法构造5，立即退出
 */
func getMaximumConsecutive(coins []int) int {
	var last int
	sort.Ints(coins)
	for _, coin := range coins {
		if coin > last + 1 {
			break
		}
		last += coin
	}
	return last + 1
}

func main() {
	fmt.Println(getMaximumConsecutive([]int{1,3}))
	fmt.Println(getMaximumConsecutive([]int{1,1,1,4}))
	fmt.Println(getMaximumConsecutive([]int{1,4,10,3,1}))
}
