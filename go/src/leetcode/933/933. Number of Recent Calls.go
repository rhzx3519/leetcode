/**
@author ZhengHao Lou
Date    2022/05/06
*/
package main

/**
https://leetcode-cn.com/problems/number-of-recent-calls/
*/

const N = 3000

type RecentCounter struct {
	req []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.req = append(this.req, t)
	i := lowerBound(this.req, t-N)
	return len(this.req) - i
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
