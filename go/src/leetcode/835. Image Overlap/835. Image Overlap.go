/**
@author ZhengHao Lou
Date    2021/11/16
*/
package main

/**
https://leetcode-cn.com/problems/image-overlap/
思路：暴力枚举
枚举偏移值，偏移范围为x方向[-n, +n], y方向[-n, +n]
time: O(N^4)
space: O(N^2)
 */
func largestOverlap(img1 [][]int, img2 [][]int) int {
	N := len(img1)
	diff := make([][]int, N<<1 + 1)
	for i := range diff {
		diff[i] = make([]int, N<<1 + 1)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if img1[i][j] == 0 {
				continue
			}

			for l := 0; l < N; l++ {
				for m := 0; m < N; m++ {
					if img2[l][m] == 0 {
						continue
					}
					// 偏移量为(i-l+N, j-m+N)
					diff[i-l+N][j-m+N]++
				}
			}
		}
	}

	var ans int
	for i := range diff {
		for j := range diff[i] {
			if diff[i][j] > ans {
				ans = diff[i][j]
			}
		}
	}

	return ans
}
