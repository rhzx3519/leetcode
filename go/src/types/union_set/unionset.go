/**
@author ZhengHao Lou
Date    2022/07/30
*/
package union_set

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
