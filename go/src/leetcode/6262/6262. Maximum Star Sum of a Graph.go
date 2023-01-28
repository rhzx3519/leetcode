/*
*
@author ZhengHao Lou
Date    2022/12/11
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

/*
*
https://leetcode.cn/problems/maximum-star-sum-of-a-graph/
*/
func maxStarSum(vals []int, edges [][]int, k int) (tot int) {
	tot = math.MinInt32
	n := len(vals)
	var g = make([][]int, n)
	var sums = make([]int, n)
	for _, e := range edges {
		a, b := e[0], e[1]
		if vals[b] > 0 {
			j := sort.Search(len(g[a]), func(i int) bool {
				return vals[g[a][i]] < vals[b]
			})
			g[a] = append(g[a], 0)
			copy(g[a][j+1:], g[a][j:])
			g[a][j] = b
			sums[a] += vals[b]
			if len(g[a]) > k {
				sums[a] -= vals[g[a][len(g[a])-1]]
				g[a] = g[a][:len(g[a])-1]
			}
		}
		
		if vals[a] > 0 {
			j := sort.Search(len(g[b]), func(i int) bool {
				return vals[g[b][i]] < vals[a]
			})
			g[b] = append(g[b], 0)
			copy(g[b][j+1:], g[b][j:])
			g[b][j] = a
			sums[b] += vals[a]
			if len(g[b]) > k {
				sums[b] -= vals[g[b][len(g[b])-1]]
				g[b] = g[b][:len(g[b])-1]
			}
		}
		
	}
	
	for i := 0; i < n; i++ {
		if sums[i]+vals[i] > tot {
			tot = sums[i] + vals[i]
		}
	}
	
	fmt.Print(g, sums, tot)
	return
}

func main() {
	maxStarSum([]int{1, 2, 3, 4, 10, -10, -20}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {3, 5}, {3, 6}}, 2)
}
