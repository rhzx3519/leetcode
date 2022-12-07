/**
@author ZhengHao Lou
Date    2022/10/18
*/
package main

import "strconv"

/**
https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/solution/zui-da-wei-n-de-shu-zi-zu-he-by-leetcode-f3yi/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func atMostNGivenDigitSet(digits []string, n int) int {
	m := len(digits)
	s := strconv.Itoa(n)
	k := len(s)
	dp := make([][2]int, k+1)
	dp[0][1] = 1
	for i := 1; i <= k; i++ {
		for _, d := range digits {
			if d[0] == s[i-1] {
				dp[i][1] = dp[i-1][1]
			} else if d[0] < s[i-1] {
				dp[i][0] += dp[i-1][1]
			} else {
				break
			}
		}
		if i > 1 {
			dp[i][0] += m + dp[i-1][0]*m
		}
	}
	return dp[k][0] + dp[k][1]
}
