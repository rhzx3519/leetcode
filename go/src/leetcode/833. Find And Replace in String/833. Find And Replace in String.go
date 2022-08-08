/**
@author ZhengHao Lou
Date    2021/11/16
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-and-replace-in-string/
思路：记录下标
time: O(n)
space: O(n)
 */
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	n := len(s)
	r := make([]int, n)
	for i := range r {
		r[i] = -1
	}

	bytes := []byte(s)
	for k := range indices {
		i := indices[k]
		if string(bytes[i:i+len(sources[k])]) == sources[k] {
			r[i] = k
		}
	}

	bytes = []byte{}
	for i := 0; i < n; i++ {
		if r[i] == -1 {
			bytes = append(bytes, s[i])
			continue
		}
		k := r[i]
		i += len(sources[k]) - 1
		bytes = append(bytes, []byte(targets[k])...)
	}

	fmt.Println(r)
	fmt.Println(string(bytes))
	return string(bytes)
}

func main() {
	findReplaceString("abcd", []int{0,2}, []string{"a","cd"}, []string{"eee","ffff"})
	findReplaceString("abcd", []int{0,2}, []string{"ab","ec"}, []string{"eee","ffff"})
}


