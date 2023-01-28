/**
@author ZhengHao Lou
Date    2023/01/09
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/
*/
func reinitializePermutation(n int) (tot int) {
	var arr, perm = make([]int, n), make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	
	var equal = func() bool {
		for i := range perm {
			if perm[i] != i {
				return false
			}
		}
		return true
	}
	
	for {
		for i := range perm {
			if i&1 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		copy(perm, arr)
		fmt.Println(perm, arr)
		tot++
		if equal() {
			return
		}
		
	}
	
}

func main() {
	fmt.Println(reinitializePermutation(6))
}
