/**
@author ZhengHao Lou
Date    2022/01/30
*/
package main

import (
	"fmt"
)

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	n := len(s)
	var hash int
	for i := 0; i < n; i++ {
		if i < k {
			hash = (hash + ((int(s[i]-'a') + 1) * pow(power, i, modulo)))
		} else {
			hash = hash + (int(s[i]-'a')+1)*pow(power, k-1, modulo)
			hash -= (int(s[i-k]-'a') + 1)
			hash /= power
		}

		if i >= k-1 {
			fmt.Println(s[i-k+1:i+1], hash)
			if hash%modulo == hashValue {
				fmt.Println(s[i-k+1 : i+1])
				return s[i-k+1 : i+1]
			}

		}
	}
	return "-1"
}

func pow(x, n, mod int) int {
	var res = 1
	var x_contribute = x
	for n > 0 {
		if n&1 == 1 {
			res = (res * x_contribute) % mod
		}
		x_contribute *= x_contribute
		n >>= 1
	}
	return res
}

func main() {
	//subStrHash("leetcode", 7, 20, 2, 0)
	//subStrHash("fbxzaad", 31, 100, 3, 32)
	//subStrHash("xmmhdakfursinye", 96, 45, 15, 21)
	//"nekv"
	subStrHash("xxterzixjqrghqyeketqeynekvqhc", 15, 94, 4, 16)
}
