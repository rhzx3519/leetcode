/**
@author ZhengHao Lou
Date    2021/10/03
*/
package main

/**
思路：
1. 不能选择让 所有已移除石子 的价值 总和 可以被 3 整除的石子
2. 尽可能让对手选择 所有已移除石子 的价值 总和 可以被 3 整除 的石子
 */
func stoneGameIX(stones []int) bool {
	var c = []int{0,0,0}
	for _, stone := range stones {
		c[stone%3]++
	}

	// c[0]可以让奇偶顺序改变，例如如果c[0]是偶数，则play1,2都可以在对手选择c[0]后，自己也选择c[0]，让奇偶序列不变
	if c[0]&1 == 0 {
		return c[1] != 0 && c[2] != 0
	}
	return abs(c[1] - c[2]) > 2
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
