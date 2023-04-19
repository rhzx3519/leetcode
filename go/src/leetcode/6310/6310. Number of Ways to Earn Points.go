package main

import "fmt"

/*
*
https://leetcode.cn/problems/number-of-ways-to-earn-points/
思路：01背包
*/
const mod int = 1e9 + 7

func waysToReachTarget(target int, types [][]int) int {
	f := make([]int, target+1)
	f[0] = 1

	for _, t := range types {
		count, marks := t[0], t[1]
		for s := target; s > 0; s-- {
			for k := 1; k <= count && s-k*marks >= 0; k++ {
				f[s] += f[s-k*marks]
			}
			f[s] %= mod
		}
	}

	return f[target]
}

func main() {
	fmt.Println(waysToReachTarget(6, [][]int{{6, 1}, {3, 2}, {2, 3}}))
}
