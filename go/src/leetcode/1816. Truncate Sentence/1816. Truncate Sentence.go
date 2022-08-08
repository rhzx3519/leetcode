/**
@author ZhengHao Lou
Date    2021/12/06
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/truncate-sentence/
 */
func truncateSentence(s string, k int) string {
	words := strings.Split(s, " ")
	return strings.Join(words[:k], " ")
}
