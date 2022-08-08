/**
@author ZhengHao Lou
Date    2022/06/09
*/
package main

import "math/rand"

type Solution struct {
	rects   [][]int
	weights []int
}

func Constructor(rects [][]int) Solution {
	var weights = make([]int, len(rects)+1)
	for i := range rects {
		m := rects[i][2] - rects[i][0] + 1
		n := rects[i][3] - rects[i][1] + 1
		weights[i+1] = weights[i] + m*n
	}
	return Solution{
		weights: weights,
		rects:   rects,
	}
}

func (this *Solution) Pick() []int {
	k := rand.Intn(this.weights[len(this.weights)-1])
	i := lowerBound(this.weights, k+1) - 1
	n := this.rects[i][3] - this.rects[i][1] + 1
	return []int{this.rects[i][0] + (k-this.weights[i])/n, this.rects[i][1] + (k-this.weights[i])%n}
}

func lowerBound(nums []int, x int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)>>1
		if nums[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
