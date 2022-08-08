/**
@author ZhengHao Lou
Date    2022/07/30
*/
package main

/**
https://leetcode.cn/problems/largest-component-size-by-common-factor/
*/
func largestComponentSize(nums []int) int {
	n := int(1e5) + 1
	uf := NewUnionFind(n)
	for i := range nums {
		for j := 2; j*j <= nums[i]; j++ {
			if nums[i]%j == 0 {
				uf.Merge(nums[i], j)
				uf.Merge(nums[i], nums[i]/j)
			}
		}
	}

	var ans int
	counter := make([]int, n+1)
	for i := range nums {
		x := uf.Find(nums[i])
		counter[x]++
		ans = max(ans, counter[x])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type UnionFind struct {
	parent, rank []int
}

func NewUnionFind(n int) UnionFind {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return UnionFind{
		parent: parent,
		rank:   make([]int, n),
	}
}
func (uf UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf UnionFind) Merge(x, y int) {
	x, y = uf.Find(x), uf.Find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
	} else {
		uf.parent[y] = x
		uf.rank[x]++
	}
}
