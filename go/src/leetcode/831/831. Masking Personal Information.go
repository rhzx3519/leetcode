package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
*
https://leetcode.cn/problems/masking-personal-information/
*/
func maskPII(s string) (ans string) {
	if strings.IndexByte(s, '@') != -1 { // email
		s = strings.ToLower(s)
		strs := strings.Split(s, "@")
		ans = fmt.Sprintf("%v*****%v@%v", strs[0][0:1], strs[0][len(strs[0])-1:], strs[1])
	} else {
		var digits []byte
		for i := range s {
			if unicode.IsDigit(rune(s[i])) {
				digits = append(digits, s[i])
			}
		}
		var nation []byte
		n := len(digits)
		for i := 0; i < n-10; i++ {
			nation = append(nation, '*')
		}
		if len(nation) == 0 {
			ans = fmt.Sprintf("***-***-%v", string(digits[n-4:]))
		} else {
			ans = fmt.Sprintf("+%v-***-***-%v", string(nation), string(digits[n-4:]))
		}

	}
	return ans
}

func main() {
	fmt.Println(maskPII("LeetCode@LeetCode.com"))
	fmt.Println(maskPII("AB@qq.com"))
	fmt.Println(maskPII("1(234)567-890"))
	fmt.Println(maskPII("86-(10)12345678"))
}
