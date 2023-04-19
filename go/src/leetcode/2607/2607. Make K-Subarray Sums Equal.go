package main

import "sort"

/*
*
https://leetcode.cn/problems/make-k-subarray-sums-equal/
a[i] = a[i+k]
a[i+1] = a[i+k+1]
根据mod(k)分组，同组内的数要求相同，
同组内都变为中数位的代价最小
*/
func makeSubKSumEqual(arr []int, k int) (tot int64) {
	n := len(arr)
	k = gcd(n, k) // k
	g := make([][]int, k)
	for i, x := range arr {
		g[i%k] = append(g[i%k], x)
	}
	for _, b := range g {
		sort.Ints(b)
		for i := range b {
			tot += int64(abs(b[i] - b[len(b)/2]))
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
