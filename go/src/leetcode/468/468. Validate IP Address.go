/**
@author ZhengHao Lou
Date    2022/05/29
*/
package main

import (
	"fmt"
	"strings"
	"unicode"
)

/**
https://leetcode.cn/problems/validate-ip-address/
*/
func validIPAddress(queryIP string) string {
	if strings.IndexByte(queryIP, '.') != -1 {
		nums := strings.Split(queryIP, ".")
		if len(nums) != 4 {
			return "Neither"
		}
		for _, num := range nums {
			if !isIPv4(num) {
				return "Neither"
			}
		}
		return "IPv4"
	} else if strings.IndexByte(queryIP, ':') != -1 {
		nums := strings.Split(queryIP, ":")
		if len(nums) != 8 {
			return "Neither"
		}
		for _, num := range nums {
			if !isIPv6(num) {
				return "Neither"
			}
		}
		return "IPv6"
	}
	return "Neither"
}

func isIPv6(s string) bool {
	n := len(s)
	if n < 1 || n > 4 {
		return false
	}
	for _, ch := range s {
		if !(unicode.IsDigit(ch) || (ch >= 'a' && ch <= 'f') ||
			(ch >= 'A' && ch <= 'F')) {
			return false
		}
	}
	return true
}

func isIPv4(s string) bool {
	n := len(s)
	if n == 0 || (s[0] == '0' && n > 1) {
		return false
	}

	var x int
	for i := 0; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
		x = x*10 + int(s[i]-'0')
		if x > 255 {
			return false
		}
	}

	return true
}

func main() {
	//fmt.Println(validIPAddress("172.16.254.1"))
	//fmt.Println(validIPAddress("2001:0db8:85a3:0:0:8A2E:0370:7334"))
	//fmt.Println(validIPAddress("256.256.256.256"))
	fmt.Println(validIPAddress("192.0.0.1"))
}
