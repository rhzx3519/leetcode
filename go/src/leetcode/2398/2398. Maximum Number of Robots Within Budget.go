/**
@author ZhengHao Lou
Date    2022/09/06
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/maximum-number-of-robots-within-budget/
https://leetcode.cn/problems/maximum-number-of-robots-within-budget/solution/by-endlesscheng-7ukp/
思路：单调栈 + 滑动窗口
*/
func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	n := len(chargeTimes)
	var ans int
	var st []int
	var s int64
	var l, r int
	for ; r < n; r++ {
		for len(st) != 0 && chargeTimes[st[len(st)-1]] <= chargeTimes[r] {
			st = st[:len(st)-1]
		}
		st = append(st, r)
		s += int64(runningCosts[r])
		for len(st) > 0 && int64(chargeTimes[st[0]])+s*int64(r-l+1) > budget {
			if st[0] == l {
				st = st[1:]
			}
			s -= int64(runningCosts[l])
			l++
		}
		ans = max(ans, r-l+1)
	}
	fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maximumRobots([]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25)
	maximumRobots([]int{11, 12, 19}, []int{10, 8, 7}, 19)
}
