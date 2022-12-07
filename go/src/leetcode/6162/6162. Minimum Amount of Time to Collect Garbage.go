/**
@author ZhengHao Lou
Date    2022/08/29
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/minimum-amount-of-time-to-collect-garbage/
*/
func garbageCollection(garbage []string, travel []int) int {
	var cost int
	//n := len(garbage)
	var ms, ps, gs int
	for i := range garbage {
		var m, p, g int
		for j := range garbage[i] {
			switch garbage[i][j] {
			case 'M':
				m++
			case 'P':
				p++
			case 'G':
				g++
			}
		}
		cost += m + p + g

		if m != 0 {
			ms = i
		}
		if p != 0 {
			ps = i
		}
		if g != 0 {
			gs = i
		}
	}

	for i, t := range travel {
		if i < ms {
			cost += t
		}
		if i < ps {
			cost += t
		}
		if i < gs {
			cost += t
		}
	}

	fmt.Println(ms, ps, gs, cost)
	return cost
}

func main() {
	garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3})
	garbageCollection([]string{"MMM", "PGM", "GP"}, []int{3, 10})
}
