/**
@author ZhengHao Lou
Date    2022/02/21
*/
package main

/**
https://leetcode-cn.com/problems/maximum-split-of-positive-even-integers/
思路：贪心
*/
func maximumEvenSplit(finalSum int64) []int64 {
	var ans = []int64{}
	if finalSum&1 == 1 {
		return ans
	}
	for i := int64(2); i <= finalSum; i += 2 {
		ans = append(ans, i)
		finalSum -= i
	}
	if finalSum > 0 {
		ans = append(ans[:len(ans)-1], finalSum+ans[len(ans)-1])
	}
	return ans
}
