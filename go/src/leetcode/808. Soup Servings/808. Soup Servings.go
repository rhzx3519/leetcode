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
	var N = n/25
	if n%25 != 0 {
		N++
	}

	mem := make(map[[2]int]float64)
	var dfs func(a, b int) float64
	dfs = func(a, b int) float64 {
		k := [2]int{a,b}
		if p, ok := mem[k]; ok {
			return p
		}
		if a <= 0 && b <= 0 {
			return 0.5
		}
		if a <= 0 {
			return 1.0
		}
		if b <= 0 {
			return 0.0
		}

		var t = 0.25*(dfs(a-4, b) + dfs(a-3, b-1) + dfs(a-2, b-2) + dfs(a-1, b-3))
		mem[k] = t
		return t
	}

	return dfs(N, N)
}

func main() {
	fmt.Println(soupServings(500*25))
}
