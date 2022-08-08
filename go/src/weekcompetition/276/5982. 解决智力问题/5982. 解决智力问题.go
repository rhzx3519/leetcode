/**
@author ZhengHao Lou
Date    2022/01/16
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/solving-questions-with-brainpower/
思路：DP，倒装遍历
f[i]表示解决[i:]问题可以获得最大分数
f[i] = f[i+1] if 不解决i问题
f[i] = pointsi + f[i + brainpoweri + 1] if 解决i问题
 */
func mostPoints(questions [][]int) int64 {
	n := len(questions)
	f := make([]int, n + 1)
	for i := n-1; i >= 0; i-- {
		point, brainpower := questions[i][0], questions[i][1]
		f[i] = max(f[i+1], f[min(i + brainpower + 1, n)] + point)
	}

	return int64(f[0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mostPoints([][]int{{3,2},{4,3},{4,4},{2,5}}))
	fmt.Println(mostPoints([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}}))
}