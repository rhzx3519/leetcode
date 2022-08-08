/**
@author ZhengHao Lou
Date    2021/12/03
*/
package main

/**
https://leetcode-cn.com/problems/greatest-sum-divisible-by-three/
思路：动态规划
f[0,1,2]分别表示元素和模为0、1、2的最大元素和
 */
func maxSumDivThree(nums []int) int {
	r := [3]int{}
	var a, b, c int
	for _, num := range nums {
		a = r[0] + num
		b = r[1] + num
		c = r[2] + num
		r[a%3] = max(r[a%3], a)
		r[b%3] = max(r[b%3], b)
		r[c%3] = max(r[c%3], c)
	}
	//fmt.Println(a, b, c)
	return r[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxSumDivThree([]int{3,6,5,1,8})
}