/**
@author ZhengHao Lou
Date    2022/01/01
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/convert-1d-array-into-2d-array/
 */
func construct2DArray(original []int, m int, n int) [][]int {
	l := len(original)
	if l != m*n {
		return [][]int{}
	}
	ans := make([][]int, 0)
	for i := 0; i < m; i++ {
		ans = append(ans, []int{})
		for j := 0; j < n; j++ {
			k := i*n + j
			ans[i] = append(ans[i], original[k])
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	construct2DArray([]int{1,2,3,4}, 2, 2)
}
