/**
@author ZhengHao Lou
Date    2022/07/18
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/contain-virus/
*/
var offset = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func containVirus(isInfected [][]int) int {
	m, n := len(isInfected), len(isInfected[0])

	var encode = func(x, y int) int {
		return x*n + y + 1
	}
	var decode = func(k int) (x, y int) {
		k -= 1
		return k / n, k % n
	}

	var bfs func(k int) (wall int, infected map[int]int)
	bfs = func(k int) (wall int, infected map[int]int) {
		x, y := decode(k)
		if isInfected[x][y] != 1 {
			return
		}
		infected = map[int]int{}
		isInfected[x][y] = -k

		var que = []int{encode(x, y)}
		for len(que) != 0 {
			cx, cy := decode(que[0])
			que = que[1:]
			for _, dxy := range offset {
				nx, ny := cx+dxy[0], cy+dxy[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if isInfected[nx][ny] == 1 {
						isInfected[nx][ny] = -k
						que = append(que, encode(nx, ny))
					} else if isInfected[nx][ny] == 0 {
						wall++
						infected[encode(nx, ny)]++
					}
				}
			}
		}
		return
	}

	var ans int

	for {
		var mx, mw, mk = 0, 0, -1
		var infectedMap = map[int]map[int]int{}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				k := encode(i, j)
				wall, infected := bfs(k)
				infectedMap[k] = infected

				if len(infected) > mx {
					mx = len(infected)
					mw = wall
					mk = k
				}
			}
		}
		if mk == -1 {
			break
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] < 0 {
					if isInfected[i][j] == -mk {
						isInfected[i][j] = 2
					} else {
						isInfected[i][j] = 1
					}
				}
			}
		}

		delete(infectedMap, mk)
		for _, inf := range infectedMap {
			for i := range inf {
				x, y := decode(i)
				isInfected[x][y] = 1
			}
		}

		ans += mw
	}

	fmt.Println(ans)
	return ans
}

func main() {
	containVirus([][]int{{0, 1, 0, 0, 0, 0, 0, 1}, {0, 1, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 0}})
	containVirus([][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}})
	containVirus([][]int{{1, 1, 1, 0, 0, 0, 0, 0, 0}, {1, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 0, 0, 0, 0}})
}
