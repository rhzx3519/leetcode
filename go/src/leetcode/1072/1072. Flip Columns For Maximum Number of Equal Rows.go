package main

/*
*
https://leetcode.cn/problems/flip-columns-for-maximum-number-of-equal-rows/description/
*/
func maxEqualRowsAfterFlips(matrix [][]int) int {
	patterns := make(map[string]int)
	n := len(matrix[0])
	for i := range matrix {
		bytes := []byte{'0'}
		for j := 1; j < n; j++ {
			k := bytes[len(bytes)-1]
			if matrix[i][j] != matrix[i][j-1] {
				bytes = append(bytes, byte(^int(k-'0')))
			} else {
				bytes = append(bytes, k)
			}
		}
		patterns[string(bytes)]++
	}

	var ans int
	for _, c := range patterns {
		if c > ans {
			ans = c
		}
	}
	return ans
}
