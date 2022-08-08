/**
@author ZhengHao Lou
Date    2021/12/03
*/
package 最短路径算法

/**
Floyd算法
对每个点对，尝试经过每个点进行优化，写法简单
 */
func Floyd(grid [][]int, n, k int) []int {
	for t := 0; t < n; t++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				grid[i][j] = min(grid[i][j], grid[i][k] + grid[k][j])
			}
		}
	}
	return grid[k]
}