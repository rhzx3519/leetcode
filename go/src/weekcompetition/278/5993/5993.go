/**
@author ZhengHao Lou
Date    2022/01/30
*/
package main

func findFinalValue(nums []int, original int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	for m[original] != 0 {
		original <<= 1
	}
	return original
}
