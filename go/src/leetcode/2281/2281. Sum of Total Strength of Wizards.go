/**
@author ZhengHao Lou
Date    2022/05/23
*/
package main

import (
	"fmt"
)

/**
https://leetcode.cn/problems/sum-of-total-strength-of-wizards/
*/
const MOD int = 1e9 + 7

func totalStrength(strength []int) int {
	n := len(strength)
	var ss, left, right = make([]int, n+2), make([]int, n), make([]int, n)
	var s int
	for i := 0; i < n; i++ {
		s += strength[i]
		ss[i+2] = (ss[i+1] + s) % MOD
		left[i] = -1
		right[i] = n
	}

	var st []int
	for i := n - 1; i >= 0; i-- {
		for len(st) != 0 && strength[st[len(st)-1]] > strength[i] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			left[j] = i
		}
		st = append(st, i)
	}

	st = []int{}
	for i := 0; i < n; i++ {
		for len(st) != 0 && strength[st[len(st)-1]] >= strength[i] {
			j := st[len(st)-1]
			st = st[:len(st)-1]
			right[j] = i
		}
		st = append(st, i)
	}

	fmt.Println(ss)
	fmt.Println(left)
	fmt.Println(right)

	var ans int
	for i := 0; i < n; i++ {
		l, r := left[i]+1, right[i]-1 // [l,r] 左闭右闭
		tot := ((i-l+1)*(ss[r+2]-ss[i+1]) - (r-i+1)*(ss[i+1]-ss[l])) % MOD
		ans = (ans + strength[i]*tot) % MOD
	}

	return (ans + MOD) % MOD
}

func main() {
	totalStrength([]int{1, 3, 1, 2})
	totalStrength([]int{5, 4, 6})
}
