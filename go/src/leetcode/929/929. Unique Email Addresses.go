/**
@author ZhengHao Lou
Date    2022/06/04
*/
package main

import (
	"fmt"
	"strings"
)

/**\
https://leetcode.cn/problems/unique-email-addresses/
*/
func numUniqueEmails(emails []string) int {
	var counter = map[string]int{}
	for _, email := range emails {
		strs := strings.Split(email, "@")
		local, domain := strs[0], strs[1]
		var lb, db []byte
	NEXT:
		for i := range local {
			switch local[i] {
			case '.':
			case '+':
				break NEXT
			default:
				lb = append(lb, local[i])
			}
		}
		for i := range domain {
			switch domain[i] {
			default:
				db = append(db, domain[i])
			}
		}
		counter[fmt.Sprintf("%v@%v", string(lb), string(db))]++
	}
	fmt.Println(counter)
	return len(counter)
}

func main() {
	numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"})
}
