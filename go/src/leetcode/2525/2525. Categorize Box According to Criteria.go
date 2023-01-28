/**
@author ZhengHao Lou
Date    2023/01/08
*/
package main

/**
https://leetcode.cn/problems/categorize-box-according-to-criteria/
*/
const (
	A int = 1e4
	B int = 1e9
)

func categorizeBox(length int, width int, height int, mass int) string {
	var bulky, heavy bool
	if length >= A || width >= A || height >= A || length*width*height >= B {
		bulky = true
	}
	if mass >= 100 {
		heavy = true
	}
	if bulky && heavy {
		return "Both"
	}
	if !bulky && !heavy {
		return "Neither"
	}
	if bulky {
		return "Bulky"
	}
	return "Heavy"
}
