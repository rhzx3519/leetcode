/**
@author ZhengHao Lou
Date    2023/01/10
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/cracking-the-safe/description/
*/
func crackSafe(n int, k int) string {
	highest := int(math.Pow10(n - 1))
	var ans string
	var vis = make(map[int]bool)
	var dfs func(int)
	dfs = func(node int) {
		for i := 0; i < k; i++ {
			x := node*10 + i
			if !vis[x] {
				vis[x] = true
				dfs(x % highest)
				ans += fmt.Sprintf("%v", i)
			}
		}
	}
	
	dfs(0)
	for i := 1; i < n; i++ {
		ans += "0"
	}
	return ans
}

func main() {
	fmt.Println(crackSafe(1, 2))
	fmt.Println(crackSafe(2, 2))
}
