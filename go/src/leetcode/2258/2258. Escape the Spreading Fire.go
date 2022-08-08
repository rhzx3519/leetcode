/**
@author ZhengHao Lou
Date    2022/05/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/escape-the-spreading-fire/
思路：二分 + BFS
*/
type pair struct{ x, y int }

var dxy = []pair{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}}

func maximumMinutes(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	
	var burn func(t int, fires []pair) bool
	var spreadfires func(grid [][]int, fires []pair) []pair
	
	burn = func(t int, fires []pair) bool {
		g := make([][]int, m)
		vis := make([][]bool, m)
		
		for i := range grid {
			vis[i] = make([]bool, n)
			g[i] = make([]int, n)
			copy(g[i], grid[i])
		}
		
		for ; t != 0 && len(fires) != 0; t-- {
			fires = spreadfires(g, fires)
		}
		fmt.Println(g)
		
		if g[0][0] == 1 {
			return true
		}
		que := []pair{{x: 0, y: 0}}
		vis[0][0] = true
		
		for len(que) != 0 {
			fires = spreadfires(g, fires)
			recent := map[pair]bool{}
			for _, fire := range fires {
				recent[fire] = true
			}
			for s := len(que) - 1; s >= 0; s-- {
				p := que[0]
				que = que[1:]
				
				for _, d := range dxy {
					ni, nj := p.x+d.x, p.y+d.y
					// 注意，如果你到达安全屋后，火马上到了安全屋，这视为你能够安全到达安全屋。
					if ni >= 0 && ni < m && nj >= 0 && nj < n && !vis[ni][nj] {
						if ni == m-1 && nj == n-1 && (g[ni][nj] == 0 || recent[pair{x: ni, y: nj}]) {
							return false
						}
						if g[ni][nj] == 0 {
							vis[ni][nj] = true
							que = append(que, pair{x: ni, y: nj})
						}
					}
				}
			}
		}
		
		return true
	}
	
	spreadfires = func(g [][]int, fires []pair) []pair {
		var f []pair
		for i := range fires {
			for _, d := range dxy {
				ni, nj := fires[i].x+d.x, fires[i].y+d.y
				if ni >= 0 && ni < m && nj >= 0 && nj < n && g[ni][nj] == 0 {
					f = append(f, pair{x: ni, y: nj})
					g[ni][nj] = 1
				}
			}
		}
		return f
	}
	
	var fires []pair
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				fires = append(fires, pair{x: i, y: j})
			}
		}
	}
	
	b := burn(5, fires)
	fmt.Println(b)
	l, r := 0, m*n+1
	for l < r {
		t := l + (r-l)>>1
		if burn(t, fires) {
			r = t
		} else {
			l = t + 1
		}
	}
	
	if l <= m*n {
		return l - 1
	}
	
	return 1e9
}

func main() {
	//maximumMinutes([][]int{{0, 2, 0, 0, 0, 0, 0}, {0, 0, 0, 2, 2, 1, 0}, {0, 2, 0, 0, 1, 2, 0}, {0, 0, 2, 2, 2, 0, 2}, {0, 0, 0, 0, 0, 0, 0}})
	//maximumMinutes([][]int{{0, 0, 0, 0}, {0, 1, 2, 0}, {0, 2, 0, 0}})
	//maximumMinutes([][]int{{0, 2, 0, 0, 1}, {0, 2, 0, 2, 2}, {0, 2, 0, 0, 0}, {0, 0, 2, 2, 0}, {0, 0, 0, 0, 0}})
	maximumMinutes([][]int{{0, 0, 0}, {2, 2, 0}, {1, 2, 0}})
}
