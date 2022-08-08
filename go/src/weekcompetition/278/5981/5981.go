/**
@author ZhengHao Lou
Date    2022/01/30
*/
package main

import "fmt"

func maxScoreIndices(nums []int) []int {
	n := len(nums)
	var n1 int
	for _, d := range nums {
		if d == 1 {
			n1++
		}
	}
	var ans = []int{0}
	var mx = n1
	var c int
	for i := 1; i <= n; i++ {
		if nums[i-1] == 1 {
			c++
		}
		l := i - c
		r := n1 - c
		if l+r > mx {
			mx = l + r
			ans = []int{i}
		} else if l+r == mx {
			ans = append(ans, i)
		}
	}
	fmt.Println(ans)
	return ans
}

func main() {
	maxScoreIndices([]int{0, 0, 1, 0})
	maxScoreIndices([]int{0, 0, 0})
	maxScoreIndices([]int{1, 1})
}
