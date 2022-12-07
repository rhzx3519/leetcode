/**
@author ZhengHao Lou
Date    2022/09/18
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/making-a-large-island/
*/
var dir = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func largestIsland(grid [][]int) (ans int) {
	n := len(grid)
	uf := NewUnionFind(n * n)
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	sz := make([]int, n*n)
	var dfs func(x, y int) int
	dfs = func(x, y int) int {
		var tot = 1
		for i := range dir {
			nx, ny := x+dir[i][0], y+dir[i][1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 1 && !vis[nx][ny] {
				uf.Merge(encode(x, y, n), encode(nx, ny, n))
				vis[nx][ny] = true
				tot += dfs(nx, ny)
			}
		}
		return tot
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 && !vis[i][j] {
				vis[i][j] = true
				sz[encode(i, j, n)] = dfs(i, j)
			}
		}
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				ans = max(ans, sz[uf.Find(encode(i, j, n))])
			} else {
				var tot int
				var counter = make(map[int]bool)
				for _, d := range dir {
					ni, nj := i+d[0], j+d[1]
					if ni >= 0 && ni < n && nj >= 0 && nj < n && grid[ni][nj] == 1 {
						p := uf.Find(encode(ni, nj, n))
						if !counter[p] {
							tot += sz[p]
						}
						counter[p] = true
					}
				}

				ans = max(ans, tot+1)
			}
		}
	}

	fmt.Println(uf.parent)
	fmt.Println(sz)
	fmt.Println(ans)
	return ans
}

func main() {
	largestIsland([][]int{{1, 0}, {0, 1}})
	largestIsland([][]int{{1, 1}, {1, 0}})
	largestIsland([][]int{{1, 1}, {1, 1}})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func encode(i, j, n int) int {
	return i*n + j
}

func decode(x, n int) (int, int) {
	return x / n, x % n
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
