/**
@author ZhengHao Lou
Date    2022/01/31
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-substring-with-given-hash-value/
思路：触发不满足求模恒等，所以需要倒序遍历
*/
func subStrHash(s string, power int, mod int, k int, hashValue int) string {
	n := len(s)
	var hash, p = 0, 1
	var sub string
	for i := n - 1; i >= 0; i-- {
		if i+k >= n {
			hash = (hash*power + encode(s[i])) % mod
			p = p * power % mod
		} else {
			hash = (hash*power + encode(s[i]) + mod - p*encode(s[i+k])%mod) % mod
		}
		if i+k <= n {
			if hash == hashValue {
				sub = s[i : i+k]
			}
		}
	}

	fmt.Println(sub)
	return sub
}

func encode(b byte) int {
	return int(b-'a') + 1
}

func main() {
	//subStrHash("leetcode", 7, 20, 2, 0)
	//subStrHash("fbxzaad", 31, 100, 3, 32)
	subStrHash("xmmhdakfursinye", 96, 45, 15, 21)
}
