/**
@author ZhengHao Lou
Date    2022/08/16
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/count-special-integers/
f[i][j]表示长度为i，mask=j的特殊整数数目
*/
func countSpecialNumbers(n int) int {
	vis := make([]bool, 10)
	
	var dfs func(num int) int
	dfs = func(num int) int {
		if num > n {
			return 0
		}
		
		var c int
		for i := 0; i < 10; i++ {
			if vis[i] {
				continue
			}
			vis[i] = true
			c += dfs(num*10 + i)
			vis[i] = false
		}
		return c + 1
	}
	
	var ans int
	for i := 1; i < 10; i++ {
		vis[i] = true
		ans += dfs(i)
		vis[i] = false
	}
	
	return ans
}

func main() {
	fmt.Println(countSpecialNumbers(20))
	fmt.Println(countSpecialNumbers(5))
	fmt.Println(countSpecialNumbers(135))
}
