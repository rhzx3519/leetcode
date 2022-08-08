/**
@author ZhengHao Lou
Date    2021/12/23
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/parallel-courses-iii/
思路：拓扑排序，动态规划
time: O(e)
space: O(n)
 */
func minimumTime(n int, relations [][]int, time []int) (ans int) {
	g := make([][]int, n)
	for i := range g {
		g[i] = make([]int, 0)
	}
	ind := make([]int, n)
	for _, e := range relations {
		u, v := e[0] - 1, e[1] - 1
		g[u] = append(g[u], v)
		ind[v]++
	}
	que := []int{}
	for i := range ind {
		if ind[i] == 0 {
			que = append(que, i)
		}
	}

	f := make([]int, n)	// f[i]表示i的前置课程中耗时最大的课程
	for len(que) != 0 {
		u := que[0]
		que = que[1:]
		f[u] += time[u]
		ans = max(ans, f[u])
		for _, v := range g[u] {
			f[v] = max(f[v], f[u])	// 更新v的前置课程中的最长时长
			ind[v]--
			if ind[v] == 0 {
				que = append(que, v)
			}
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTime(3, [][]int{{1,3},{2,3}}, []int{3,2,5}))
	fmt.Println(minimumTime(5, [][]int{{1,5},{2,5},{3,5},{3,4},{4,5}}, []int{1,2,3,4,5}))
}