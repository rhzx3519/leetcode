package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-number-of-teams/
思路：统计i左右两侧比rating[i]大、或者比rating[i]小的个数
time: O(n^2)
space: O(1)
 */
func numTeams(rating []int) int {
	var unit int
	n := len(rating)
	for j, m := range rating {
		var l1, l2, r1, r2 int

		for i := 0; i < j; i++ {
			if rating[i] < m {
				l1++
			} else if rating[i] > m {
				l2++
			}
		}

		for k := j+1; k < n; k ++ {
			if rating[k] < m {
				r1++
			} else if rating[k] > m {
				r2++
			}
		}

		unit += l1*r2 + l2*r1
	}
	fmt.Println(unit)
	return unit
}

func main() {
	numTeams([]int{2,5,3,4,1})
	numTeams([]int{2,1,3})
	numTeams([]int{1,2,3,4})
}
