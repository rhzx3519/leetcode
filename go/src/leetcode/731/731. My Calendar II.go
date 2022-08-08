/**
@author ZhengHao Lou
Date    2022/07/19
*/
package main

type MyCalendarTwo struct {
	overlaps [][]int
	calendar [][]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for i := range this.overlaps {
		a, b := this.overlaps[i][0], this.overlaps[i][1]
		if start < b && end > a {
			return false
		}
	}

	for i := range this.calendar {
		a, b := this.calendar[i][0], this.calendar[i][1]
		if start < b && end > a {
			this.overlaps = append(this.overlaps, []int{max(start, a), min(end, b)})
		}
	}

	this.calendar = append(this.calendar, []int{start, end})
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
