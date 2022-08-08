/**
@author ZhengHao Lou
Date    2021/12/23
*/
package main

import "sort"

/**
https://leetcode-cn.com/problems/my-calendar-iii/
思路：差分数组
t 为日程的最大时间
time: O(t)
space: O(n)
 */
type MyCalendarThree struct {
	diff map[int]int
}


func Constructor() MyCalendarThree {
	return MyCalendarThree{
		diff: make(map[int]int),
	}
}


func (this *MyCalendarThree) Book(start int, end int) (ans int) {
	this.diff[start]++
	this.diff[end]--
	keys := []int{}
	for k := range this.diff {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	var m int
	for _, k := range keys {
		m += this.diff[k]
		ans = max(ans, m)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
