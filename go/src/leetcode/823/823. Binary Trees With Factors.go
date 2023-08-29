package main

import "sort"

/*
*
https://leetcode.cn/problems/binary-trees-with-factors/description/
*/
const MOD int = 1e9 + 7

func numFactoredBinaryTrees(arr []int) (tot int) {
	f := make(map[int]int)
	d := make(map[int]int)
	for _, i := range arr {
		d[i]++
		f[i] = 1
	}
	sort.Ints(arr)
	for i, cur := range arr {
		for j := 0; j < i; j++ {
			if cur%arr[j] == 0 && d[cur/arr[j]] != 0 {
				f[cur] = (f[cur] + f[arr[j]]*f[cur/arr[j]]) % MOD
			}
		}
		tot = (tot + f[cur]) % MOD
	}
	return
}
