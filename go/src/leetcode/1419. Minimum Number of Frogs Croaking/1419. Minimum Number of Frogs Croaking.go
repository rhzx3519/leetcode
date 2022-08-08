/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

/**
https://leetcode-cn.com/problems/minimum-number-of-frogs-croaking/
 */
func minNumberOfFrogs(croakOfFrogs string) int {
	var frog int
	var c, r, o, a, k int
	for i := range croakOfFrogs {
		switch croakOfFrogs[i] {
		case 'c':
			c++
		case 'r':
			r++
		case 'o':
			o++
		case 'a':
			a++
		case 'k':
			k++
		}
		frog = max(frog, c - k)
		if c >= r && r >= o && o >= a && a >= k	{
			continue
		}
		return -1
	}

	if c == r && r == o && o == a && a == k {
		return frog
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
