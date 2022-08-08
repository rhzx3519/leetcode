/**
@author ZhengHao Lou
Date    2022/07/05
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/number-of-people-aware-of-a-secret/
思路：动态规划
f[i]: 表示第i天知道秘密的人数
inc[i]: 表示第i天分享秘密的人数
dec[i]: 表示第i天忘掉秘密的人数

第i天知道秘密的人数f[i] = f[i-1] - dec[i] + inc[i]
第i天分享密码的人数inc[i] = f[i-delay], i-delay > 0
第i天忘掉密码的人数dec[i] = f[i-forget], i-forget > 0
*/
const MOD int = 1e9 + 7

func peopleAwareOfSecret(n int, delay int, forget int) int {
	f, inc, dec := make([]int, n+1), make([]int, n+1), make([]int, n+1)
	f[1] = 1
	for i := 2; i <= n; i++ {
		if i-delay > 0 {
			inc[i] = f[i-delay]
		}
		if i-forget > 0 {
			dec[i] = f[i-forget]
		}
		f[i] = (f[i-1] - dec[i] + inc[i] + MOD) % MOD
	}

	c := (f[n] - dec[n] + MOD) % MOD
	fmt.Println(c)
	return c
}

func main() {
	peopleAwareOfSecret(6, 2, 4)
	peopleAwareOfSecret(4, 1, 3)
}
