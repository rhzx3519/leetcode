/**
@author ZhengHao Lou
Date    2021/12/15
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/loud-and-rich/
思路：拓扑排序
 */
const INF int = math.MaxInt32 >> 1
func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)

	ind := make([]int, n)
	edges := map[int][]int{}
	for _, r := range richer {
		a, b := r[0], r[1] 		// a > b
		if _, ok := edges[a]; !ok {
			edges[a] = []int{}
		}
		edges[a] = append(edges[a], b)
		ind[b]++
	}

	quietToPerson := make(map[int]int)
	que := [][]int{}	// inque nodes of 0 in degree
	for i := 0; i < n; i++ {
		if ind[i] == 0 {
			que = append(que, []int{i, quiet[i]})
		}
		quietToPerson[quiet[i]] = i
	}

	var ans = make([]int, n)
	for i := range ans {
		ans[i] = quiet[i]
	}

	for len(que) != 0 {
		p, q := que[0][0], que[0][1]
		que = que[1:]
		for _, np := range edges[p] {
			ans[np] = min(ans[np], q)
			ind[np]--
			if ind[np] == 0 {
				que = append(que, []int{np, ans[np]})
			}
		}
	}

	for i := range ans {
		ans[i] = quietToPerson[ans[i]]
	}

	fmt.Println(ans)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	loudAndRich([][]int{{1,0},{2,1},{3,1},{3,7},{4,3},{5,3},{6,3}}, []int{3,2,5,4,6,1,7,0})
	loudAndRich([][]int{}, []int{0})
}