/**
@author ZhengHao Lou
Date    2022/09/18
*/
package main

import "sort"

/**
https://leetcode.cn/problems/maximum-matching-of-players-with-trainers/
*/
func matchPlayersAndTrainers(players []int, trainers []int) int {
	var c int
	sort.Ints(players)
	sort.Ints(trainers)
	for i := range players {
		j := lowerBound(trainers, players[i])
		if j == len(trainers) {
			break
		}
		copy(trainers[j:], trainers[j+1:])
		trainers = trainers[:len(trainers)-1]
		c++
	}
	return c
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

func main() {
	matchPlayersAndTrainers([]int{1, 1000000000}, []int{1000000000, 1})
}
