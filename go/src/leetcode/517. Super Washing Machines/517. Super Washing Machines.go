/**
  @author ZhengHao Lou
  Date    2021/09/29

https://leetcode-cn.com/problems/super-washing-machines/
*/
package main

func findMinMoves(machines []int) int {
	n := len(machines)
	var s int
	for i := range machines {
		s += machines[i]
	}
	if s % n != 0 {
		return -1
	}
	avg := s / n
	s = 0
	var ans int
	for _, cloth := range machines {
		cloth -= avg
		s += cloth
		ans = max(ans, max(cloth, abs(s)))
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}