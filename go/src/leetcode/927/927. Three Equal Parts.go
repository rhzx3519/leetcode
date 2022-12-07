/**
@author ZhengHao Lou
Date    2022/10/06
*/
package main

/**
https://leetcode.cn/problems/three-equal-parts/
*/
func threeEqualParts(arr []int) []int {
	n := len(arr)
	var sum int
	for _, v := range arr {
		sum += v
	}
	if sum%3 != 0 {
		return []int{-1, -1}
	}
	if sum == 0 {
		return []int{0, 2}
	}
	var partial = sum / 3
	var cur, first, second, third int
	for i, v := range arr {
		if v == 1 {
			if cur == 0 {
				first = i
			} else if cur == partial {
				second = i
			} else if cur == partial<<1 {
				third = i
			}
			cur++
		}
	}
	length := n - third
	if first+length <= second && second+length <= third {
		for i := 0; third+i < n; i++ {
			if arr[first+i] != arr[second+i] || arr[second+i] != arr[third+i] {
				return []int{-1, -1}
			}
		}
		return []int{first + length - 1, second + length}
	}

	return []int{-1, -1}
}
