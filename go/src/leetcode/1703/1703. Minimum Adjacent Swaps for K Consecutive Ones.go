/*
*
@author ZhengHao Lou
Date    2022/12/18
*/
package main

import (
	"fmt"
	"math"
)

/*
*
https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/

作者：灵茶山艾府
链接：https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/solutions/2024387/tu-jie-zhuan-huan-cheng-zhong-wei-shu-ta-iz4v/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func minMoves(nums []int, k int) int {
	p := []int{}
	for i, v := range nums {
		if v != 0 {
			p = append(p, i-len(p))
		}
	}
	m := len(p)
	s := make([]int, m+1) // p 的前缀和
	for i, v := range p {
		s[i+1] = s[i] + v
	}
	ans := math.MaxInt
	for i, v := range s[:m-k+1] { // p[i] 到 p[i+k-1] 中所有数到 p[i+k/2] 的距离之和，取最小值
		ans = min(ans, v+s[i+k]-s[i+k/2]*2-p[i+k/2]*(k%2))
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	fmt.Println(minMoves([]int{1, 0, 0, 1, 0, 1}, 2))
	fmt.Println(minMoves([]int{1, 0, 0, 0, 0, 0, 1, 1}, 3))
	fmt.Println(minMoves([]int{1, 1, 0, 1}, 2))
}
