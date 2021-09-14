package main

import (
	"strings"
)

/**
https://leetcode-cn.com/problems/text-justification/
 */
func fullJustify(words []string, maxWidth int) []string {
	article := []string{}
	n := len(words)
	var i int

	for i < n {
		var row = []string{}
		var j = i
		var length int
		for j < n && length + len(words[j]) <= maxWidth {
			row = append(row, words[j])
			length += len(words[j]) + 1
			j++
		}
		if j < n {
			article = append(article, format(words[i:j], maxWidth))
			//fmt.Println(format(words[i:j], maxWidth))
		} else {
			article = append(article, formatLast(words[i:j], maxWidth))
			//fmt.Println(formatLast(words[i:j], maxWidth))
		}
		i = j
	}

	return article
}

func format(words []string, maxWidth int) string {
	//fmt.Println(words)
	n := len(words)
	if n == 1 {
		bytes := []byte(words[0])
		for i := len(words[0]); i < maxWidth; i++ {
			bytes = append(bytes, ' ')
		}
		return string(bytes)
	}

	var sz int
	for i := range words {
		sz += len(words[i])
	}
	var left = maxWidth - sz
	t := left / (n - 1)
	r := left % (n - 1)
	bytes := []byte{}
	for i := 0; i < n-1; i++ {
		bytes = append(bytes, []byte(words[i])...)
		for j := 0; j < t; j++ {
			bytes = append(bytes, ' ')
		}
		if r != 0 {
			bytes = append(bytes, ' ')
			r--
		}
	}
	bytes = append(bytes, []byte(words[n-1])...)

	return string(bytes)
}

func formatLast(words []string, maxWidth int) string {
	bytes := []byte(strings.Join(words, " "))
	for i := len(bytes); i < maxWidth; i++ {
		bytes = append(bytes, ' ')
	}
	return string(bytes)
}

func main() {
	//fullJustify([]string{"This", "is", "an", "example", "of", "text", "justification."}, 16)
	//fullJustify([]string{"What","must","be","acknowledgment","shall","be"}, 16)
	fullJustify([]string{"Science","is","what","we","understand","well","enough","to","explain",
		"to","a","computer.","Art","is","everything","else","we","do"}, 20)
}