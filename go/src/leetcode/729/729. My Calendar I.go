/**
@author ZhengHao Lou
Date    2022/07/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/my-calendar-i/
*/
type MyCalendar struct {
	intervals [][]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	i := bound(this.intervals, start)
	if i != len(this.intervals) && this.intervals[i][0] < end {
		return false
	}
	this.intervals = append(this.intervals, []int{})
	copy(this.intervals[i+1:], this.intervals[i:])
	this.intervals[i] = []int{start, end}
	return true
}

func bound(intervals [][]int, x int) int {
	l, r := 0, len(intervals)
	for l < r {
		m := l + (r-l)>>1
		if intervals[m][1] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Book(10, 20))
	fmt.Println(obj.Book(15, 25))
	fmt.Println(obj.Book(20, 30))
	fmt.Println(obj.intervals)
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
