package main

import (
	"fmt"
	"math"
)

/**
思路：状态压缩DP
f[mask]表示当任务安排为状态mask时，最少的工作时间段
f[0] = 0; f[1<<n -1]为最终结果.

细节

枚举 mask 的子集有一个经典的小技巧，对应的伪代码如下：
```code
subset = mask
while subset != 0 do
    // subset 是 mask 的一个子集，可以用其进行状态转移
    ...
    // 使用按位与运算在 O(1) 的时间快速得到下一个（即更小的）mask 的子集
    subset = (subset - 1) & mask
end while

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/minimum-number-of-work-sessions-to-finish-the-tasks/solution/wan-cheng-ren-wu-de-zui-shao-gong-zuo-sh-tl0p/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
```
 */
func minSessions(tasks []int, sessionTime int) int {
	n := len(tasks)
	valid := make([]bool, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		var s int
		for d := 0; d < n; d++ {
			if mask& (1<<d) != 0 {
				s += tasks[d]
			}
		}
		if s <= sessionTime {
			valid[mask] = true
		}
	}

	f := make([]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		f[i] = math.MaxInt32>>1
	}
	f[0] = 0

	for mask := 1; mask < 1<<n; mask++ {
		for subset := mask; subset != 0; subset = (subset - 1)&mask {
			if valid[subset] {
				f[mask] = min(f[mask], f[mask^subset] + 1)
			}
		}
	}

	fmt.Println(f[1<<n-1])
	return f[1<<n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minSessions([]int{1,2,3}, 3)
	minSessions([]int{3,1,3,1,1}, 8)
	minSessions([]int{1,2,3,4,5}, 15)
	minSessions([]int{1,5,7,10,3,8,4,2,6,2}, 10)
}
