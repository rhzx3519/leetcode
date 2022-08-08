/**
@author ZhengHao Lou
Date    2021/11/25
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/flower-planting-with-no-adjacent/
 */
func gardenNoAdj(n int, paths [][]int) []int {
	adj := make([][]int, n+1)
	for i := range adj {
		adj[i] = make([]int, 0)
	}
	for _, path := range paths {
		adj[path[0]] = append(adj[path[0]], path[1])
		adj[path[1]] = append(adj[path[1]], path[0])
	}

	gardens := make([]int, n + 1)
	for i := 1; i <= n; i++ {
		if gardens[i] != 0 {
			continue
		}
		flowers := make([]bool, 5)
		for _, ng := range adj[i] {
			if gardens[ng] != 0 {
				flowers[gardens[ng]] = true
				continue
			}
		}
		for k := 1; k <= 4; k++ {
			if flowers[k] == false {
				flowers[k] = true
				gardens[i] = k
				break
			}
		}
	}
	fmt.Println(gardens)
	return gardens[1:]
}

func main() {
	gardenNoAdj(3, [][]int{{1,2},{2,3},{3,1}})
	gardenNoAdj(4, [][]int{{1,2},{3,4}})
	gardenNoAdj(4, [][]int{{1,2},{2,3},{3,4},{4,1},{1,3}})
}