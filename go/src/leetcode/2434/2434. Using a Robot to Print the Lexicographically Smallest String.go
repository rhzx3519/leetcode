/**
@author ZhengHao Lou
Date    2022/10/10
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/using-a-robot-to-print-the-lexicographically-smallest-string/
*/
func robotWithString(s string) string {
	n := len(s)
	left := make([]int, n)
	var l = n - 1
	var smallest = s[n-1]
	for i := n - 1; i >= 0; i-- {
		if s[i] < smallest {
			smallest = s[i]
			l = i
		}
		left[i] = l
	}
	var bytes []byte
	var bs []byte
	for i := 0; i < n; i++ {
		if len(bs) == 0 || bs[len(bs)-1] > s[left[i]] {
			if i == left[i] {
				bytes = append(bytes, s[i])
			} else {
				bs = append(bs, s[i])
			}
		} else if len(bs) != 0 && bs[len(bs)-1] <= s[left[i]] {
			bytes = append(bytes, bs[len(bs)-1])
			bs = bs[:len(bs)-1]
			i--
		} else {
			bs = append(bs, s[i])
		}
	}

	for len(bs) != 0 {
		bytes = append(bytes, bs[len(bs)-1])
		bs = bs[:len(bs)-1]
	}

	fmt.Println(left, string(bytes))
	return string(bytes)
}

func main() {
	robotWithString("zza")
	robotWithString("bac")
	robotWithString("bdda")

	robotWithString("vzhofnpo") // "fnohopzv"
}
