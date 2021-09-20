package main

import "fmt"

/**
https://leetcode-cn.com/problems/beautiful-array/solution/932-piao-liang-shu-zu-fen-zhi-si-xiang-g-1xxg/
A[k]*2 = A[i] + A[j], i < k < j
漂亮数组已经过线性变换后，仍然是漂亮数组，即A是漂亮数组 k*A + b也是漂亮数组
time：递归调用logn次，每次递归需要遍历N长度的数组，NlogN
space: 递归调用logn次，递归栈大小是N，NlogN
 */
func beautifulArray(n int) []int {
	mem := map[int][]int{1: {1}}

	var dp func(int) []int
	dp = func(N int) []int {
		if a, ok := mem[N]; ok {
			return a
		}

		left := dp((N + 1) / 2)
		right := dp(N / 2)
		a := []int{}
		for _, x := range left {
			a = append(a, 2 * x - 1)
		}
		for _, x := range right {
			a = append(a, 2 * x)
		}
		mem[N] = a
		return a
	}

	return dp(n)
}

func main() {
	fmt.Println(beautifulArray(4))
	fmt.Println(beautifulArray(5))
}


