/**
@author ZhengHao Lou
Date    2023/01/16
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode.cn/problems/sentence-similarity-iii/
*/
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	words1, words2 := strings.Split(sentence1, " "), strings.Split(sentence2, " ")
	if len(words1) < len(words2) {
		return areSentencesSimilar(sentence2, sentence1)
	}
	var i int
	for ; i < len(words2) && words1[i] == words2[i]; i++ {
	}
	if i == len(words2) {
		return true
	}
	var j, k = len(words1) - 1, len(words2) - 1
	for ; k >= 0 && words1[j] == words2[k]; {
		j--
		k--
	}
	if k == -1 {
		return true
	}
	if k < i {
		return true
	}
	
	return false
}

func main() {
	//fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	//fmt.Println(areSentencesSimilar("of", "A lot of words"))
	//fmt.Println(areSentencesSimilar("Eating right now", "Eating"))
	//fmt.Println(areSentencesSimilar("Luky", "Lucccky"))
	fmt.Println(areSentencesSimilar("A B C D B B", "A B B"))
}
