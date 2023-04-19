package main

import "fmt"

/*
*
https://leetcode.cn/problems/split-the-array-to-make-coprime-products/
思路：
1. 两个数互质 等价于 两个数没有相同的质因数
2. 有效分割点满足，质因数p的范围在[i,j]之间时，有效分割点必不在[i,j]之间
本质：跳跃游戏
*/
func findValidSplit(nums []int) int {
	n := len(nums)
	left := make(map[int]int) // 记录质因数p第一次出现的下标
	right := make([]int, n)   // 记录左端点为i的右端点最大值

	f := func(p, i int) {
		if j, ok := left[p]; !ok {
			left[p] = i
		} else {
			right[j] = i
		}
	}

	for i, x := range nums {
		for d := 2; d*d <= x; d++ {
			if x%d == 0 {
				f(d, i)
				for x /= d; x%d == 0; x /= d {
				}
			}
		}
		if x > 1 {
			f(x, i)
		}
	}

	var maxR int
	for i := 0; i < n; i++ {
		if i > maxR {
			return i - 1
		}
		maxR = max(maxR, right[i])
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findValidSplit([]int{4, 7, 8, 15, 3, 5}))
	fmt.Println(findValidSplit([]int{4, 7, 15, 8, 3, 5}))
}
