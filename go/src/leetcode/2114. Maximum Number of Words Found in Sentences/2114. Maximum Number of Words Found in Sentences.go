/**
@author ZhengHao Lou
Date    2022/01/04
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/maximum-number-of-words-found-in-sentences/
 */
func mostWordsFound(sentences []string) int {
	var mi int
	for _, s := range sentences {
		t := strings.Split(s, " ")
		if len(t) > mi {
			mi = len(t)
		}
	}
	return mi
}
