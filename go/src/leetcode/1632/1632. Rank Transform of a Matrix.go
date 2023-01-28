/**
@author ZhengHao Lou
Date    2023/01/25
*/
package main

import "sort"

/**
https://leetcode.cn/problems/rank-transform-of-a-matrix/description/
*/
func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	var encode = func(x, y int) int {
		return x*n + y
	}
	var decode = func(x int) (int, int) {
		return x / n, x % n
	}
	
	counter := make(map[int][]int)
	for i := range matrix {
		for j := range matrix[i] {
			x := matrix[i][j]
			counter[x] = append(counter[x], encode(i, j))
		}
	}
	rowMax := make([]int, m)
	colMax := make([]int, n)
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	
	var vals []int
	for v := range counter {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	for _, v := range vals {
		ps := counter[v]
		ds := NewDisjointSet(m + n)
		rank := make([]int, m+n)
		for _, p := range ps {
			i, j := decode(p)
			ds.Union(i, j+m)
		}
		for _, p := range ps {
			i, j := decode(p)
			rank[ds.Find(i)] = max(rank[ds.Find(i)], max(rowMax[i], colMax[j]))
		}
		for _, p := range ps {
			i, j := decode(p)
			ans[i][j] = 1 + rank[ds.Find(i)]
			rowMax[i], colMax[j] = ans[i][j], ans[i][j]
		}
	}
	
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type DisjoinSet struct {
	n      int
	parent []int
	size   []int
}

func (d DisjoinSet) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d DisjoinSet) Union(x, y int) bool {
	px, py := d.Find(x), d.Find(y)
	if px != py {
		if d.size[px] > d.size[py] {
			d.parent[py] = px
			d.size[px] += d.size[py]
		} else {
			d.parent[px] = py
			d.size[py] += d.size[px]
		}
		return true
	}
	return false
}

func (d DisjoinSet) Size(x int) int {
	return d.size[d.Find(x)]
}

func NewDisjointSet(n int) DisjoinSet {
	parent := make([]int, n)
	sz := make([]int, n)
	for i := range parent {
		parent[i] = i
		sz[i] = 1
	}
	return DisjoinSet{
		n:      n,
		parent: parent,
		size:   sz,
	}
}
