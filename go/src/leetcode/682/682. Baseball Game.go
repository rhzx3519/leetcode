/**
@author ZhengHao Lou
Date    2022/03/26
*/
package main

import "strconv"

/**
https://leetcode-cn.com/problems/baseball-game/
*/
func calPoints(ops []string) int {
	var st []int
	for _, op := range ops {
		switch op {
		case "+":
			st = append(st, st[len(st)-1]+st[len(st)-2])
		case "D":
			st = append(st, st[len(st)-1]<<1)
		case "C":
			st = st[:len(st)-1]
		default:
			i, _ := strconv.Atoi(op)
			st = append(st, i)
		}
	}
	var s int
	for i := range st {
		s += st[i]
	}
	return s
}

func main() {
	calPoints([]string{"5", "2", "C", "D", "+"})
}
