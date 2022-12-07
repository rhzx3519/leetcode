/**
@author ZhengHao Lou
Date    2022/10/05
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
https://leetcode.cn/problems/subdomain-visit-count/
*/
func subdomainVisits(cpdomains []string) (tot []string) {
	counter := make(map[string]int)
	for _, d := range cpdomains {
		tmp := strings.Split(d, " ")
		c, _ := strconv.Atoi(tmp[0])
		strs := strings.Split(tmp[1], ".")
		for i := range strs {
			counter[strings.Join(strs[i:], ".")] += c
		}
	}

	for i := range counter {
		tot = append(tot, fmt.Sprintf("%v %v", counter[i], i))
	}
	return tot
}
