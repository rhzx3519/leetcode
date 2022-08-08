/**
@author ZhengHao Lou
Date    2021/11/18
*/
package main

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
		d.parent[py] = px
		d.size[px] += d.size[py]
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
