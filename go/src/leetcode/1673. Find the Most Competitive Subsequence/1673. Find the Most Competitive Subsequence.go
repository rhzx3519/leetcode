/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-the-most-competitive-subsequence/
思路：等价于求长度为n - k的最小数字, 单调栈
 */
func mostCompetitive(nums []int, k int) []int {
	k = len(nums) - k
	st := []int{}
	for _, num := range nums {
		for k > 0 && len(st) > 0 && st[len(st) - 1] > num {
			k--
			st = st[:len(st) - 1]
		}
		st = append(st, num)
	}
	st = st[:len(st) - k]
	return st
}

func main() {
	fmt.Println(mostCompetitive([]int{3,5,2,6}, 2))
	fmt.Println(mostCompetitive([]int{2,4,3,3,5,4,9,6}, 4))
}
