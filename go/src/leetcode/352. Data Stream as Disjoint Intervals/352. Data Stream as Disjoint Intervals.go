/**
@author ZhengHao Lou
Date    2021/10/09
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/data-stream-as-disjoint-intervals/
 */
type SummaryRanges struct {
	intervals [][]int
	arr []int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		intervals: [][]int{},
		arr: []int{},
	}
}


func (this *SummaryRanges) AddNum(val int)  {
	i := lowerBound(this.arr, val)
	if i != len(this.arr) && this.arr[i] == val {
		return
	}

	this.arr = append(this.arr, 0)
	copy(this.arr[i+1:], this.arr[i:])
	this.arr[i] = val
}


func (this *SummaryRanges) GetIntervals() [][]int {
	n := len(this.arr)
	intervals := [][]int{}
	var last int
	for i := 0; i < n - 1; i++ {
		if this.arr[i+1] != this.arr[i] + 1 {
			intervals = append(intervals, []int{this.arr[last], this.arr[i]})
			last = i + 1
		}
	}
	intervals = append(intervals, []int{this.arr[last], this.arr[n-1]})
	fmt.Println(this.arr, intervals)
	return intervals
}

func lowerBound(arr []int, x int) int {
	l, r := 0, len(arr)
	var m int
	for l < r {
		m = l + (r- l)>>1
		if arr[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	obj := Constructor()
	obj.AddNum(6)
	obj.GetIntervals()
	obj.AddNum(6)
	obj.GetIntervals()
	obj.AddNum(7)
	obj.GetIntervals()
	obj.AddNum(2)
	obj.GetIntervals()
	obj.AddNum(6)
	obj.GetIntervals()
}


/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */
