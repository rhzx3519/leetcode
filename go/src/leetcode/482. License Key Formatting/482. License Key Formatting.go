/**
@author ZhengHao Lou
Date    2021/10/04
*/
package main

import "strings"

/**
https://leetcode-cn.com/problems/license-key-formatting/
 */
func licenseKeyFormatting(s string, k int) string {
	s = strings.ReplaceAll(strings.ToUpper(s), "-", "")
	n := len(s)
	r := n % k
	strs := []string{}
	if r > 0 {
		strs = append(strs, s[:r])
	}
	for i := r; i < n; i += k {
		strs = append(strs, s[i:i+k])
	}
	return strings.Join(strs, "-")
}

func main() {
	licenseKeyFormatting("5F3Z-2e-9-w", 4)
}