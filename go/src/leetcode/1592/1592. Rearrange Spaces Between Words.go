/**
@author ZhengHao Lou
Date    2022/09/07
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode.cn/problems/rearrange-spaces-between-words/
*/
func reorderSpaces(text string) string {
	n := len(text)
	var words []string
	for _, s := range strings.Split(text, " ") {
		if s != "" {
			words = append(words, s)
		}
	}
	var nw, nb int
	for _, w := range words {
		nw += len(w)
	}
	nb = n - nw
	var c, r = 0, nb
	if len(words) > 1 {
		c = nb / (len(words) - 1)
		r = nb % (len(words) - 1)
	}
	var sc, sr []byte
	for i := 0; i < c; i++ {
		sc = append(sc, ' ')
	}
	for i := 0; i < r; i++ {
		sr = append(sr, ' ')
	}
	return strings.Join(words, string(sc)) + string(sr)
}

func main() {
	fmt.Println(reorderSpaces("  this   "))
}
