/**
@author ZhengHao Lou
Date    2022/01/06
*/
package main

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/simplify-path/
 */
func simplifyPath(path string) string {
	ps := strings.Split(path, "/")
	//fmt.Println(ps)
	st := []string{}
	for _, s := range ps {
		switch s {
		case "":
		case "/":
		case ".":
		case "..":
			if len(st) > 0 {
				st = st[:len(st)-1]
			}
		default:
			st = append(st, s)
		}
	}
	//fmt.Println(st)
	return "/" + strings.Join(st,  "/")
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
