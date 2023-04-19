/**
@author ZhengHao Lou
Date    2022/01/04
*/
package main

/**
https://leetcode-cn.com/problems/cat-and-mouse/
https://leetcode-cn.com/problems/cat-and-mouse/solution/gong-shui-san-xie-dong-tai-gui-hua-yun-y-0bx1/
 */
const N int = 55
var (
	f = [2*N*N][N][N]int{}
)

func catMouseGame(graph [][]int) int {
	n := len(graph)

	for k := 0; k < 2*n*n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				f[k][i][j] = -1
			}
		}
	}

	// 0: draw / 1: mouse / 2: cat
	var dfs func(k, a, b int) int
	dfs = func(k, a, b int) int {
		ans := f[k][a][b]
		if a == 0 {
			ans = 1
		} else if a == b {
			ans = 2
		} else if k >= 2*n*n {
			ans = 0
		} else if ans == -1 {
			if k%2 == 0 { // mouse move
				var win, draw bool
				for _, ne := range graph[a] {
					t := dfs(k+1, ne, b)
					if t == 1 {
						win = true
					} else if t == 0 {
						draw = true
					}
					if win {
						break
					}
				}
				if win {
					ans = 1
				} else if draw {
					ans = 0
				} else {
					ans = 2
				}

			} else { // cat move
				var win, draw bool
				for _, ne := range graph[b] {
					if ne == 0 {
						continue
					}
					t := dfs(k+1, a, ne)
					if t == 2 {
						win = true
					} else if t == 0 {
						draw = true
					}
					if win {
						break
					}
				}

				if win {
					ans = 2
				} else if draw {
					ans = 0
				} else {
					ans = 1
				}
			}
		}

		f[k][a][b] = ans
		return ans
	}


	return dfs(0,  1, 2)
}
