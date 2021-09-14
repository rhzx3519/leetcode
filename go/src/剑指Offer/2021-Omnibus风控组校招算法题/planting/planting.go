package main

import "fmt"

/**
2.(Medium)
小张正在玩涂色游戏，他将a*b张相同的白纸片摆成矩阵，利用调色盘(共有n种颜料)给白纸片涂满对应颜色，
由于调色盘上颜料有限，每种颜料只够给 colors[i] 张白纸片涂色，且所有颜料刚好能涂满全部白纸片，
每张白纸片只能涂一种颜料，小张希望每张白纸片的相邻上下左右的纸片的颜色不能与这个纸片相同，
请问是否能按他的想法完成涂色呢？(保证用尽所有颜料)

- 1≤x,y≤5
- ∑c[i] = x*y
输出"true" 或者 "false"
 */

var offset = [][]int{{-1,0}, {0,-1}}

func planting(m, n int, colors []int) bool {
	var total = m * n
	dp := make([]int, total)
	for i := range dp {
		dp[i] = -1
	}

	var backtrace func(idx int) bool
	backtrace = func(idx int) bool {
		if idx == total {
			//for i := 0; i < m; i++ {
			//	fmt.Println(dp[i*n:(i+1)*n])
			//}
			return true
		}
		x, y := idx/n, idx%n

		used := []int{}
		for _, dxy := range offset {
			nx := x + dxy[0]
			ny := y + dxy[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			used = append(used, dp[nx*n + ny])
		}

		for i := range colors {
			if colors[i] == 0 {
				continue
			}
			if in(i, used) {
				continue
			}
			colors[i]--
			dp[idx] = i
			if backtrace(idx+1) {
				return true
			}
			dp[idx] = -1
			colors[i]++
		}

		return false
	}

	return backtrace(0)
}

func in(c int, used []int) bool {
	for _, color := range used {
		if c == color {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(planting(1, 5, []int{4,1})) 		// false
	fmt.Println(planting(1, 1, []int{1})) 		// true
	fmt.Println(planting(2, 2, []int{1,2,1}))	// true
	fmt.Println(planting(5, 5, []int{5,14,1,5}))	// false
	fmt.Println(planting(5, 4, []int{3,3,11,3}))	// false
	fmt.Println(planting(2, 4, []int{2,3,3}))	// true
	fmt.Println(planting(5, 5, []int{12,12,1}))	// true
}
