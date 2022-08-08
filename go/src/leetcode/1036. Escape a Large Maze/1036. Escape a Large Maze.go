/**
@author ZhengHao Lou
Date    2022/01/11
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/escape-a-large-maze/
 */
const (
	BASE int = 131
	N int = 1e6
)

var offset = [][]int{{1,0},{0,-1},{-1,0},{0,1}}

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	blockedMap := make(map[int]bool)
	for _, b := range blocked {
		blockedMap[b[0]*BASE + b[1]] = true
	}
	n := len(blockedMap)
	MAX := n*(n-1)/2
	return bfs(source, target, blockedMap, MAX) && bfs(target, source, blockedMap, MAX)

}

func bfs(s, t []int, blockedMap map[int]bool, MAX int) bool {
	vis := make(map[int]bool)
	vis[hash(s)] = true
	que := [][]int{s}
	for len(que) != 0 && len(vis) <= MAX {
		x, y := que[0][0], que[0][1]
		que = que[1:]
		for _, dxy := range offset {
			nx, ny := x + dxy[0], y + dxy[1]
			if nx < 0 || nx >= N || ny < 0 || ny >= N {
				continue
			}
			if nx == t[0] && ny == t[1] {
				return true
			}
			z := nx*BASE + ny
			if blockedMap[z] || vis[z] {
				continue
			}
			que = append(que, []int{nx, ny})
			vis[z] = true
		}
	}
	return len(vis) > MAX
}

func hash(c []int) int {
	return c[0]*BASE + c[1]
}

func main() {
	fmt.Println(isEscapePossible([][]int{{0,1},{1,0}}, []int{0,0}, []int{0,2}))
	fmt.Println(isEscapePossible([][]int{}, []int{0,0}, []int{999999,999999}))
}
