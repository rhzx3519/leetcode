/**
@author ZhengHao Lou
Date    2022/09/15
*/
package main

/**
https://leetcode.cn/problems/bulb-switcher-ii/
*/
func flipLights(n int, presses int) int {
	if presses == 0 {
		return 1
	}
	switch n {
	case 1:
		return 2
	case 2:
		if presses == 1 {
			return 3
		}
		return 4
	default:
		if presses == 1 {
			return 4
		} else if presses == 2 {
			return 7
		}
		return 8
	}
}
