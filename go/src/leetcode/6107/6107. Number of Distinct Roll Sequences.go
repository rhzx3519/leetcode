/**
@author ZhengHao Lou
Date    2022/06/27
*/
package main

/**
https://leetcode.cn/problems/number-of-distinct-roll-sequences/
定义 f[i][\textit{last}][\textit{last}_2]f[i][last][last
2
​
 ] 表示序列长度为 ii，最后一个元素是 \textit{last}last，倒数第二个元素是 \textit{last}_2last
2
​
  的序列数目。

通过枚举 \textit{last}last 和 \textit{last}_2last
2
​
 ，我们可以计算出 f[i+1][j][\textit{last}]f[i+1][j][last]，需满足

作者：endlesscheng
链接：https://leetcode.cn/problems/number-of-distinct-roll-sequences/solution/by-endlesscheng-tgkn/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/
const (
	N   int = 1e4
	MOD int = 1e9 + 7
)

var f = [N + 1][6][6]int{}

func init() {
	for last := 0; last < 6; last++ {
		for last2 := 0; last2 < 6; last2++ {
			if last2 != last && gcd(last2+1, last+1) == 1 {
				f[2][last2][last] = 1
			}
		}
	}
	for i := 2; i < N; i++ {
		for j := 0; j < 6; j++ {
			for last := 0; last < 6; last++ {
				if last != j && gcd(last+1, j+1) == 1 {
					for last2 := 0; last2 < 6; last2++ {
						if last2 != j {
							f[i+1][j][last] = (f[i+1][j][last] + f[i][last][last2]) % MOD
						}
					}
				}
			}
		}
	}
}

func distinctSequences(n int) (ans int) {
	if n == 1 {
		return 6
	}
	for _, row := range f[n] {
		for _, v := range row {
			ans = (ans + v) % MOD
		}
	}
	return
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
