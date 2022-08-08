/**
@author ZhengHao Lou
Date    2021/10/30
*/
package main

import "fmt"

/**
https://leetcode-cn.com/contest/biweekly-contest-64/problems/kth-distinct-string-in-an-array/
 */
func kthDistinct(arr []string, k int) string {
	n := len(arr)
	t := make([]bool, n)
	mapper := make(map[string]int)
	for i, s := range arr {
		if j, ok := mapper[s]; ok {
			t[i] = true
			t[j] = true
		}
		mapper[s] = i
	}

	fmt.Println(t)
	var index int
	for i := range arr {
		if t[i] == false {
			index ++
		}
		if index == k {
			return arr[i]
		}
	}
	return ""
}

func main() {
	fmt.Println(kthDistinct([]string{"d","b","c","b","c","a"}, 2))
	fmt.Println(kthDistinct([]string{"aaa","aa","a"}, 1))
	fmt.Println(kthDistinct([]string{"a","b","a"}, 3))
}
