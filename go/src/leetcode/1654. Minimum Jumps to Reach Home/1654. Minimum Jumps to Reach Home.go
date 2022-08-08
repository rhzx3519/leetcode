/**
@author ZhengHao Lou
Date    2021/11/23
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-jumps-to-reach-home/
解题思路
由于我们可以超出家的位置，最短路算法可能超时，故我们需要减小搜索范围。
可以证明，一定可以在下标 [0,M + a + b][0,M+a+b] 的范围内找到最优解，M = max(F_{i},x)M=max(Fi,x)，F_{i}
  是各个禁止点，xx 是家的位置。因为 M,a,b <= 2000，也就是说搜索范围不会超过 6000。

作者：newhar
链接：https://leetcode-cn.com/problems/minimum-jumps-to-reach-home/solution/dao-jia-de-zui-shao-tiao-yue-ci-shu-zui-duan-lu-zh/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
const (
	M = 6001
	INF = math.MaxInt32 >> 1
)
func minimumJumps(forbidden []int, a int, b int, x int) int {
	f := make([]int, M)
	for i := range f {
		f[i] = INF
	}
	for _, i := range forbidden {
		f[i] = -1
	}
	var dfs func(cur, step int, turn bool)
	dfs = func(cur, step int, turn bool) {
		if cur < 0 || cur >= M || step >= f[cur] {
			return
		}
		f[cur] = step
		if turn {
			dfs(cur -  b, step + 1, false)
		}
		dfs(cur + a, step + 1, true)
	}

	dfs(0, 0, false)
	if f[x] == INF {
		return -1
	}
	return f[x]
}

func main() {
	fmt.Println(minimumJumps([]int{14,4,18,1,15}, 3, 15, 9))
}
