/**
@author ZhengHao Lou
Date    2022/11/07
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/next-greater-element-iv/
思路：单调栈
使用t保存第一次弹出的元素，当t中的元素第二次弹出时，第二大的元素=x
*/
func secondGreaterElement(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var s, t []int
	for i, x := range nums {
		for len(t) != 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x
			t = t[:len(t)-1]
		}

		j := len(s) - 1
		for ; j >= 0 && nums[s[j]] < x; j-- {
		}
		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}

	fmt.Println(ans)
	return ans
}

func main() {
	secondGreaterElement([]int{2, 4, 0, 9, 6})
	secondGreaterElement([]int{3, 3})
}
