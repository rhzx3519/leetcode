/*
*
@author ZhengHao Lou
Date    2022/12/05
*/
package main

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	n := len(words)
	for i := 0; i < n; i++ {
		if words[i][len(words[i])-1] != words[(i+1)%n][0] {
			return false
		}
	}
	return true
}
