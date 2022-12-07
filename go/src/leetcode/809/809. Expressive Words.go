/*
*
@author ZhengHao Lou
Date    2022/11/25
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/expressive-words/
*/
func expressiveWords(s string, words []string) (tot int) {
	var target [][]int

	for i := range s {
		j := int(s[i] - 'a')
		if len(target) == 0 || target[len(target)-1][0] != j {
			target = append(target, []int{j, 1})
		} else {
			target[len(target)-1][1]++
		}
	}
OUT:
	for _, word := range words {
		var k int
		var i int
		for i < len(word) {
			j := i + 1
			for j < len(word) && word[j] == word[j-1] {
				j++
			}
			l := j - i
			if k >= len(target) || target[k][0] != int(word[i]-'a') ||
				target[k][1] < l || (target[k][1] != l && target[k][1] < 3) {
				//fmt.Println(target[k], l)
				i = j
				continue OUT
			}
			i = j
			k++
		}
		if k == len(target) {
			tot++
		}
	}
	fmt.Println(tot)
	return
}

func main() {
	expressiveWords("heeellooo", []string{"hello", "hi", "helo"})
}
