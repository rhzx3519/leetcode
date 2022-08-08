/**
@author ZhengHao Lou
Date    2021/10/03
*/
package main

func minimumMoves(s string) int {
	var count int
	var i int
	for i < len(s) {
		if s[i] == 'X' {
			count++
			i += 3
		} else {
			i++
		}
	}
	return count
}