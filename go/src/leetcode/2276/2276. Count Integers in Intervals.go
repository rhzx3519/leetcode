/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

/**
https://leetcode.cn/problems/count-integers-in-intervals/
*/
type CountIntervals struct {
	intervals [][]int
	c         int
}

func Constructor() CountIntervals {
	return CountIntervals{}
}

func (this *CountIntervals) Add(l int, r int) {
	i := lowerBound(this.intervals, l) // i := lowerBound(this.intervals, l - 1), 可以合并[5,6][7,10]这种相邻的区间
	for i < len(this.intervals) {
		if this.intervals[i][0] > r { // this.intervals[i][0] > r +  1 ，可以合并[5,6][7,10]这种相邻的区间
			break
		}
		l = min(l, this.intervals[i][0])
		r = max(r, this.intervals[i][1])
		this.c -= this.intervals[i][1] - this.intervals[i][0] + 1
		copy(this.intervals[i:], this.intervals[i+1:])
		this.intervals = this.intervals[:len(this.intervals)-1]
	}
	this.c += r - l + 1
	this.intervals = append(this.intervals, []int{})
	copy(this.intervals[i+1:], this.intervals[i:])
	this.intervals[i] = []int{l, r}
}

func (this *CountIntervals) Count() int {
	return this.c
}

func lowerBound(intervals [][]int, x int) int {
	l, r := 0, len(intervals)
	for l < r {
		m := l + (r-l)>>1
		if intervals[m][1] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
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

func main() {
	c := Constructor()
	c.Add(2, 3)
	c.Add(7, 10)
	c.Count()
	c.Add(5, 8)
	c.Count()
}
