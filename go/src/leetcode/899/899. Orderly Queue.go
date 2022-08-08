/**
@author ZhengHao Lou
Date    2022/08/03
*/
package main

import "sort"

/**
https://leetcode.cn/problems/orderly-queue/
*/
func orderlyQueue(s string, k int) string {
	var bytes = []byte(s)
	if k > 1 {
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		return string(bytes)
	}

	var ms = ""
	n := len(bytes)
	for i := 0; i < n; i++ {
		s := string(bytes[i:]) + string(bytes[:i])
		if ms == "" || ms > s {
			ms = s
		}
	}

	return ms
}

func main() {
	orderlyQueue("v", 1)
}
