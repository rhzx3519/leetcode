/**
@author ZhengHao Lou
Date    2021/11/16
*/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
https://leetcode-cn.com/problems/subdomain-visit-count/
 */
func subdomainVisits(cpdomains []string) []string {
	counter := make(map[string]int)
	for _, d := range cpdomains {
		sp := strings.SplitN(d, " ", 2)
		c, _ := strconv.Atoi(sp[0])
		names := strings.Split(sp[1], ".")
		n := len(names)
		var bytes = []byte(names[n-1])
		counter[string(bytes)] += c
		for i := n - 2; i >= 0; i-- {
			bytes = append([]byte(names[i]), append([]byte{'.'}, bytes...)...)
			counter[string(bytes)] += c
		}
	}
	var ans = []string{}
	for name, c := range counter {
		ans = append(ans, fmt.Sprintf("%v %v", c, name))
	}
	return ans
}

func main() {
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
