package main

import (
	"fmt"
)

/**
https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/
思路：转化求最长递增子序列
https://leetcode-cn.com/problems/minimum-operations-to-make-a-subsequence/solution/de-dao-zi-xu-lie-de-zui-shao-cao-zuo-ci-hefgl/

 */
func minOperations(target []int, arr []int) int {
	mapper := make(map[int]int)
	for i, num := range target {
		mapper[num] = i
	}
	a := []int{}
	for _, num := range arr {
		if i, ok := mapper[num]; ok {
			a = append(a, i)
		}
	}

	d := []int{}	// 求数组a的最长递增子序列长度
	for _, i := range a {
		idx := lowerBound(d, i)
		if idx == len(d) {
			d = append(d, i)
		} else {
			d[idx] = i
		}
	}
	fmt.Println(d)
	return len(target) - len(d)
}

func lowerBound(arr []int, target int) int {
	l, r := 0, len(arr)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if arr[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}


func main() {
	fmt.Println(minOperations([]int{5,1,3}, []int{9,4,2,3,4}))
	fmt.Println(minOperations([]int{6,4,8,1,3,2}, []int{4,7,6,2,3,8,6,1}))
}