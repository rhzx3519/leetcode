/**
@author ZhengHao Lou
Date    2022/01/08
*/
package main

import "strings"

func capitalizeTitle(title string) string {
	strs := []string{}
	for _, s := range strings.Split(title, " ") {
		s = strings.ToLower(s)
		if len(s) > 2 {
			s = strings.ToUpper(string(s[0])) + s[1:]
		}
		strs = append(strs, s)
	}

	return strings.Join(strs, " ")
}
