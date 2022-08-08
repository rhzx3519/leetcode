/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/minimum-moves-to-make-array-complementary/
思路：差分数组
考虑选择的数对和为target，则对每一对数对(a, b), a<=b, 如果：
1) target < 1 + a, 需要操作2次
2) 1 + a <= target < a + b, 需要操作1次
3) target == a + b, 不要操作
4) a + b < target <= limit + b, 需要操作1次
5) limit + b < target, 需要操作2次
 */
func minMoves(nums []int, limit int) int {
	n := len(nums)
	delta := make([]int, limit<<1 + 2)
	for i := 0; i < n>>1; i++ {
		a, b := min(nums[i], nums[n - i - 1]), max(nums[i], nums[n - i - 1])
		delta[0] += 2
		delta[1 + a] -= 2
		delta[1 + a] += 1
		delta[a + b] -= 1
		delta[a + b + 1] += 1
		delta[limit + b + 1] -= 1
		delta[limit + b + 1] += 2
	}

	var ans = math.MaxInt32 >> 1
	var steps int
	for i := 0; i < limit<<1 + 2; i++ {
		steps += delta[i]
		ans = min(ans, steps)
	}
	return ans
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
	fmt.Println(minMoves([]int{1,2,4,3}, 4))
	fmt.Println(minMoves([]int{1,2,2,1}, 2))
	fmt.Println(minMoves([]int{1,2,1,2}, 2))
}
