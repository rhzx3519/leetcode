/**
@author ZhengHao Lou
Date    2022/11/02
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode.cn/problems/coordinate-with-maximum-network-quality/
*/
const N = 50

func bestCoordinate(towers [][]int, radius int) (ans []int) {
	g := make([][]int, N)
	for i := range g {
		g[i] = make([]int, N)
	}

	for i := 0; i <= N; i++ {
		for j := 0; j <= N; j++ {
			for _, tower := range towers {
				d2 := (i-tower[0])*(i-tower[0]) + (j-tower[1])*(j-tower[1])
				if d2 <= radius*radius {
					g[i][j] += int(float64(tower[2]) / float64((1 + math.Sqrt(float64(d2)))))
				}
			}
		}
	}

	ans = []int{0, 0}
	var mx int
	for i := range g {
		for j := range g[i] {
			if g[i][j] > mx {
				mx = g[i][j]
				ans = []int{i, j}
			}
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	bestCoordinate([][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}, 2)
	bestCoordinate([][]int{{23, 11, 21}}, 9)
	bestCoordinate([][]int{{1, 2, 13}, {2, 1, 7}, {0, 1, 9}}, 2)
}
