/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/detect-cycles-in-2d-grid/
 */

// 解法1：并查集
func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	dset := NewDisjointSet(m*n)
	for i := range grid {
		for j := range grid[i] {
			// 上侧格子
			if i > 0 && grid[i-1][j] == grid[i][j] {
				if !dset.Union(i*n+j, (i-1)*n+j) {
					return true
				}
			}
			// 左侧格子
			if j > 0 && grid[i][j] == grid[i][j-1] {
				if !dset.Union(i*n+j, i*n+j-1) {
					return true
				}
			}
		}
	}
	return false
}

type DisjoinSet struct {
	n int
	parent []int
}

func (d DisjoinSet) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find( d.parent[x])
	}
	return d.parent[x]
}

func (d DisjoinSet) Union(x, y int) bool {
	px, py := d.Find(x), d.Find(y)
	if px != py {
		d.parent[py] = px
		return true
	}
	return false
}

func NewDisjointSet(n int) DisjoinSet {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return DisjoinSet{
		n: n,
		parent: parent,
	}
}

// 解法2：DFS
var dxy = [][]int{{1,0},{0,-1},{-1,0},{0,1}}	// 0, 1, 2, 3
func containsCycle2(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(r, c, d int) bool
	dfs = func(r, c, d int) bool {

		for k := range dxy {
			if k == d {
				continue
			}
			nx, ny := dxy[k][0] + r, dxy[k][1] + c
			if nx < 0 || nx >= m || ny < 0 || ny >= n || grid[nx][ny] != grid[r][c] {
				continue
			}
			if vis[nx][ny] {
				return true
			}
			vis[nx][ny] = true
			if dfs(nx, ny, (k+2)%4) {
				return true
			}
		}
		return false
	}

	for i := range grid {
		for j := range grid[i] {
			if !vis[i][j] {
				vis[i][j] = true
				if dfs(i, j, -1) {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	fmt.Println(containsCycle2([][]byte{{'a','a','a','a'},{'a','b','b','a'},{'a','b','b','a'},{'a','a','a','a'}}))
	fmt.Println(containsCycle2([][]byte{{'c','c','c','a'},{'c','d','c','c'},{'c','c','e','c'},{'f','c','c','c'}}))
	fmt.Println(containsCycle2([][]byte{{'a','b','b'},{'b','z','b'},{'b','b','a'}}))
}









