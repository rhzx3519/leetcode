/**
还记得我们的定义吗？dp[i][j][len]dp[i][j][len] 表示从字符串 S 中 i 开始长度为 len 的字符串是否能变换为从字符串 T 中 j 开始长度为
len 的字符串，所以答案是 dp[0][0][n]。 时间复杂度 O(N^4)，空间复杂度O(N^3)

作者：jerry_nju
链接：https://leetcode-cn.com/problems/scramble-string/solution/miao-dong-de-qu-jian-xing-dpsi-lu-by-sha-yu-la-jia/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
package main

import "fmt"

func isScramble(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	m, n := len(s1), len(s2)
	if m !=n {
		return false
	}

	dp := make([][][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]bool, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]bool, n+1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dp[i][j][1] = s1[i]	== s2[j]
		}
	}

	// 枚举区间长度 2～n
	for l := 2; l <= n; l++ {
		// 枚举 S 中的起点位置
		for i := 0; i <= n-l; i++ {
			// 枚举 T 中的起点位置
			for j := 0; j <= n-l; j++ {
				// 枚举划分位置
				for k := 1; k < l ; k++ {
					// 第一种情况：S1 -> T1, S2 -> T2
					if dp[i][j][k] && dp[i+k][j+k][l-k] {
						dp[i][j][l] = true
						break
					}
					// 第二种情况：S1 -> T2, S2 -> T1
					// S1 起点 i，T2 起点 j + 前面那段长度 len-k ，S2 起点 i + 前面长度k
					if dp[i][j+l-k][k] && dp[i+k][j][l-k] {
						dp[i][j][l] = true
						break
					}
				}
			}
		}
	}

	return dp[0][0][n]
}

func main() {
	fmt.Println(isScramble("great", "rgeat"))
	fmt.Println(isScramble("abcde", "caebd"))
	fmt.Println(isScramble("a", "a"))
}