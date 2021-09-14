package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-all-groups-of-farmland/
 */
func findFarmland(land [][]int) [][]int {
	m, n := len(land), len(land[0])
	ans := [][]int{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if land[i][j] != 1 {
				continue
			}

			var i1 = i
			for i1 < m && land[i1][j] == 1 {
				i1++
			}

			var j1 = j
			for j1 < n && land[i][j1] == 1 {
				j1++
			}

			for x := i; x < i1; x++ {
				for y := j; y < j1; y++ {
					land[x][y] = 3
				}
			}

			ans = append(ans, []int{i,j,i1-1,j1-1})
		}
	}


	fmt.Println(ans)
	return ans
}

func main() {
	findFarmland([][]int{{1,0,0},{0,1,1},{0,1,1}})
	findFarmland([][]int{{1,1},{1,1}})
	findFarmland([][]int{{0}})
}
