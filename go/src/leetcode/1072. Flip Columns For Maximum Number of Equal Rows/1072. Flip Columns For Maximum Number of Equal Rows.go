/**
@author ZhengHao Lou
Date    2021/11/10
*/
package main

/**
https://leetcode-cn.com/problems/flip-columns-for-maximum-number-of-equal-rows/
思路：记录跳变模式相同的行的数量
 */
func maxEqualRowsAfterFlips(matrix [][]int) int {
	_, n := len(matrix), len(matrix[0])
	patterns := make(map[string]int)
	for i := range matrix {
		bytes := []byte{'0'}
		for j := 1; j < n; j++ {
			k := bytes[len(bytes) - 1]
			if matrix[i][j] != matrix[i][j-1] {
				bytes = append(bytes,  byte(^int(k - '0')))
			} else {
				bytes = append(bytes,  k)
			}
		}
		patterns[string(bytes)]++
	}

	var ans int
	for _, c := range patterns {
		ans = max(ans, c)
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

