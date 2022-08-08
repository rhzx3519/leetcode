/**
@author ZhengHao Lou
Date    2021/12/11
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/online-election/
思路：记录每个时刻票数最多的候选人
 */
type TopVotedCandidate struct {
	arr [][]int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	arr := [][]int{}
	var votes = make(map[int]int)
	var max, maxP int
	for i := 0; i < len(persons); i++ {
		t, p := times[i], persons[i]
		votes[p]++
		if votes[p] >= max {
			max = votes[p]
			maxP = p
		}
		arr = append(arr, []int{t, maxP})
	}
	return TopVotedCandidate{arr: arr}
}


func (this *TopVotedCandidate) Q(t int) int {
	i := uppperBound(this.arr, t)
	if i == 0 {
		return -1
	}
	fmt.Println(this.arr[i-1][1])
	return this.arr[i-1][1]
}

func uppperBound(arr [][]int, x int) int {
	l, r := 0, len(arr)
	var m int
	for l < r {
		m = l + (r - l) >> 1
		if arr[m][0] > x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	obj := Constructor([]int{0,1,1,0,0,1,0}, []int{0,5,10,15,20,25,30})
	obj.Q(3)
	obj.Q(12)
	obj.Q(25)
	obj.Q(15)
	obj.Q(24)
	obj.Q(8)
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
