/**
@author ZhengHao Lou
Date    2021/11/18
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/possible-bipartition/
思路：涂色问题
由相互讨厌的人构成一个图，图中每个响铃的节点必须是不同的颜色，如果有冲突，则无法分成两组
 */
func possibleBipartition(n int, dislikes [][]int) bool {
	dis := make([]map[int]int, n+1)	// 讨厌的人
	for i := 1; i <= n; i++ {
		dis[i] = map[int]int{}
	}
	for _, d := range dislikes {
		a, b := d[0], d[1]
		dis[a][b]++
		dis[b][a]++
	}

	// 相互讨厌的人涂成不同的颜色，如果有冲突，则无法分成两组
	colors :=  make([]int, n+1)
	var paint func(idx, color int) bool
	paint = func(i, c int) bool {
		if colors[i] != 0 {
			return colors[i] == c
		}
		colors[i] = c
		for k := range dis[i] {
			if !paint(k, c^1) {
				return false
			}
		}
		return true
	}

	for i := 1; i <= n; i++ {
		if colors[i] == 0 && !paint(i, 0) {
			return false
		}
	}

	return true
}


func main() {
	fmt.Println(possibleBipartition(4, [][]int{{1,2},{1,3},{2,4}}))
	fmt.Println(possibleBipartition(3, [][]int{{1,2},{1,3},{2,3}}))
	fmt.Println(possibleBipartition(5, [][]int{{1,2},{2,3},{3,4},{4,5},{1,5}}))
}
