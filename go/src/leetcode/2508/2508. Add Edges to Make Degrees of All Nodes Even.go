/*
*
@author ZhengHao Lou
Date    2022/12/19
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/add-edges-to-make-degrees-of-all-nodes-even/
思路：分类讨论
*/
func isPossible(n int, edges [][]int) bool {
	var g = make([]map[int]bool, n+1)
	for _, e := range edges {
		if g[e[0]] == nil {
			g[e[0]] = make(map[int]bool)
		}
		if g[e[1]] == nil {
			g[e[1]] = make(map[int]bool)
		}
		g[e[0]][e[1]] = true
		g[e[1]][e[0]] = true
	}

	var odd []int
	for i := 1; i <= n; i++ {
		if g[i] != nil && len(g[i])&1 == 1 {
			odd = append(odd, i)
		}
	}
	if len(odd) == 0 {
		return true
	}
	if len(odd) == 2 {
		a, b := odd[0], odd[1]
		if !g[a][b] && !g[b][a] {
			return true
		} else if g[a][b] {
			for i := 1; i <= n; i++ {
				if i != a && i != b && !g[i][a] && !g[i][b] {
					return true
				}
			}
		}
		return false
	}

	if len(odd) == 4 {
		a, b, c, d := odd[0], odd[1], odd[2], odd[3]
		return (!g[a][b] && !g[c][d]) ||
			(!g[a][c] && !g[b][d]) ||
			(!g[a][d] && !g[b][c])
	}

	return false
}

func main() {
	//fmt.Println(isPossible(4, [][]int{{4, 1}, {3, 2}, {2, 4}, {1, 3}}))
	fmt.Println(isPossible(17, [][]int{{11, 12}, {11, 8}, {2, 4}, {15, 6}, {6, 11}, {2, 11}, {12, 16}}))
}
