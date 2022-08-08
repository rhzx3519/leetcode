/**
@author ZhengHao Lou
Date    2022/05/02
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
https://leetcode-cn.com/problems/tag-validator/
*/
func isValid(code string) bool {
	var tags []string
	for code != "" {
		if code[0] != '<' {
			if len(tags) == 0 {
				return false
			}
			code = code[1:]
			continue
		}
		if len(code) == 1 {
			return false
		}
		if code[1] == '/' {
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			tag := code[2:j]
			if len(tags) == 0 || tags[len(tags)-1] != tag {
				return false
			}
			tags = tags[:len(tags)-1]
			code = code[j+1:]
			if len(tags) == 0 && code != "" {
				return false
			}
		} else if code[1] == '!' {
			if len(tags) == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
				return false
			}
			j := strings.Index(code, "]]>")
			if j == -1 {
				return false
			}
			code = code[j+1:]
		} else {
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			tag := code[1:j]
			if len(tag) == 0 || len(tag) > 9 {
				return false
			}
			for _, ch := range tag {
				if !unicode.IsUpper(ch) {
					return false
				}
			}
			tags = append(tags, tag)
			code = code[j+1:]
		}
	}
	return len(tags) == 0
}

func main() {
	fmt.Println(isValid("<DIV>This is the first line <![CDATA[<div>]]></DIV>"))
	fmt.Println(isValid("<A>  <B> </A>   </B>"))
}
