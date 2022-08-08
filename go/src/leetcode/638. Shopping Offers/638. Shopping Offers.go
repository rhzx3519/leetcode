/**
@author ZhengHao Lou
Date    2021/10/24
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/shopping-offers/
思路：大礼包 + 物品, 记忆化搜索

 */
const N int = 6
func shoppingOffers(price []int, special [][]int, needs []int) int {
	n := len(price)

	var cache = make(map[[N]int]int)
	var dfs func(neds [N]int) int
	dfs = func(neds [N]int) int {
		if p, ok := cache[neds]; ok {
			return p
		}
		var ans int
		// 不购买礼包时的成本
		for i := 0; i < n; i++ {
			ans += neds[i] * price[i]
		}
		// 遍历每个礼包
		for i := range special {
			next := [N]int{}
			for k := 0; k < N; k++ {
				next[k] = neds[k]
			}
			var valid = true
			for j := 0; j < n; j++ {
				if special[i][j] > neds[j] {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}

			for j := 0; j < n; j++ {
				next[j] -= special[i][j]
			}
			ans = min(ans, dfs(next) + special[i][n])
		}

		cache[neds] = ans
		return ans
	}

	ned := [N]int{}
	for i := range needs {
		ned[i] = needs[i]
	}
	return dfs(ned)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shoppingOffers([]int{2,5}, [][]int{{3,0,5},{1,2,10}}, []int{3,2}))
}