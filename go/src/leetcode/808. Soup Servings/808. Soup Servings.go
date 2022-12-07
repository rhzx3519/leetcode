/**
@author ZhengHao Lou
Date    2021/11/16
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/soup-servings/
*/
func soupServings(n int) float64 {
	if n >= 500*25 {
		return 1
	}
	N := n / 25
	if n%25 != 0 {
		N++
	}
	var mem = map[[2]int]float64{}
	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		if x, ok := mem[[2]int{a, b}]; ok {
			return x
		}
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1
		}
		if b <= 0 {
			return 0
		}
		x := 0.25 * (dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		mem[[2]int{a, b}] = x
		return x
	}

	return dfs(N, N)
}

func main() {
	fmt.Println(soupServings(500 * 25))
}
