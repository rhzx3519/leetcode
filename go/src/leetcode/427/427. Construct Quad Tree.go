/**
@author ZhengHao Lou
Date    2022/04/29
*/
package main

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

/**
https://leetcode-cn.com/problems/construct-quad-tree/
*/
func construct(grid [][]int) *Node {
	var dfs func(x, y, l int) *Node
	dfs = func(x, y, l int) *Node {
		for i := x; i < x+l; i++ {
			for j := y; j < y+l; j++ {
				if grid[i][j] != grid[x][y] {
					return &Node{
						Val:         true,
						IsLeaf:      false,
						TopLeft:     dfs(x, y, l>>1),
						TopRight:    dfs(x, y+l>>1, l>>1),
						BottomLeft:  dfs(x+l>>1, y, l>>1),
						BottomRight: dfs(x+l>>1, y+l>>1, l>>1),
					}
				}
			}
		}
		return &Node{
			Val:    grid[x][y] == 1,
			IsLeaf: true,
		}
	}
	
	return dfs(0, 0, len(grid))
}
