/**
@author ZhengHao Lou
Date    2021/11/26
*/
package main

import "sort"

/**
https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/
思路：贪心
1. 最大移动次数
s1=stones[n−1]−stones[0]+1−n。
s2=min(stones[n−1]−stones[n−2]−1,stones[1]−stones[0]−1)。
mx = s1 - s2
2. 最小移动次数
如果最后游戏结束，那么一定有 nn 个连续坐标摆满了石子。如果我们要移动最少，必定要找一个石子序列，使得在 nn 大小连续的坐标内，初始时有最多的石子。
设想有个尺子，上面有 nn 个刻度点，我们用这个尺子在石子从最左边到最右边移动，每动一次都查看下在尺子范围内有 mm 个石子，那么要使这个区间填满，就需要移动 n-mn−m 次。
只要在尺子外部有石子，就有策略填满尺子内的。
这些次数中最小的就为虽少次数。
但是有一种特例：
1，2，3，4，7
这种 1-4 是最好的序列，但是 7 不能移动到端点，只能 1 先移动到 6，然后 7 移动到 5 解决，这种情况要用 2 步。就是尺子内的石子都是连续的，中间没空洞，只在边上有空，要用 2 次。

链接：https://leetcode-cn.com/problems/moving-stones-until-consecutive-ii/solution/jie-ti-si-lu-by-owenzzz/
 */

func numMovesStonesII(stones []int) []int {
	n := len(stones)
	sort.Ints(stones)

	var mi, mx int
	s1 := stones[n-1] - stones[0] + 1 - n
	s2 := min(stones[n-1] - stones[n-2] - 1, stones[1] - stones[0] - 1)
	mx = s1 - s2
	mi = mx
	var i, j int
	for i = 0; i < n; i++ {
		for j+1 < n && stones[j+1] - stones[i] + 1 <= n {
			j++
		}
		var cost = n - (j - i + 1)
		if j-i+1 == n-1 && stones[j] - stones[i] + 1 == n - 1 {
			cost = 2
		}
		mi = min(mi, cost)
	}

	return []int{mi, mx}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}