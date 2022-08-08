/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/
思路：利用时间戳dfs，可以求出节点i是否是节点j的祖先节点
*/
func minimumScore(nums []int, edges [][]int) int {
	var ans = math.MaxInt32
	n := len(nums)
	g := make([][]int, n)
	for _, e := range edges {
		g[e[0]] = append(g[e[0]], e[1])
		g[e[1]] = append(g[e[1]], e[0])
	}
	in, out := make([]int, n), make([]int, n)
	xor := make([]int, n)
	var clock int

	var dfs func(x, px int)
	dfs = func(x, px int) {
		clock++
		xor[x] = nums[x]
		in[x] = clock
		for _, nx := range g[x] {
			if nx != px {
				dfs(nx, x)
				xor[x] ^= xor[nx]
			}
		}

		out[x] = clock
	}

	dfs(0, -1)
	var isAncester func(x, y int) bool
	isAncester = func(x, y int) bool {
		return in[x] < in[y] && in[y] <= out[x]
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			var a, b, c int
			if isAncester(i, j) {
				a, b, c = xor[0]^xor[i], xor[i]^xor[j], xor[j]
			} else if isAncester(j, i) {
				a, b, c = xor[0]^xor[j], xor[j]^xor[i], xor[i]
			} else {
				a, b, c = xor[0]^xor[i]^xor[j], xor[i], xor[j]
			}
			d := max(max(a, b), c) - min(a, min(b, c))
			ans = min(ans, d)
		}
	}

	fmt.Println(ans)
	return ans
}

func main() {
	minimumScore([]int{1, 5, 5, 4, 11}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}})
	minimumScore([]int{5, 5, 2, 4, 4, 2}, [][]int{{0, 1}, {1, 2}, {5, 2}, {3, 4}, {1, 3}})
	minimumScore([]int{32, 8, 11}, [][]int{{0, 1}, {1, 2}})
}

type DisjointSet struct {
	n       int
	parent  []int
	counter map[int]int
}

func NewDisjointSet(nums []int) *DisjointSet {
	n := len(nums)
	var parent = make([]int, n)
	counter := map[int]int{}
	for i := 0; i < n; i++ {
		counter[i] = nums[i]
		parent[i] = i
	}

	return &DisjointSet{
		n:       n,
		parent:  parent,
		counter: counter,
	}
}

func (d *DisjointSet) Find(x int) int {
	if d.parent[x] != x {
		return d.Find(d.parent[x])
	}
	return x
}

func (d *DisjointSet) Union(a, b int) {
	pa, pb := d.Find(a), d.Find(b)
	if pa != pb {
		d.counter[pa] ^= d.counter[pb]
		d.parent[pb] = pa
		delete(d.counter, pb)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
