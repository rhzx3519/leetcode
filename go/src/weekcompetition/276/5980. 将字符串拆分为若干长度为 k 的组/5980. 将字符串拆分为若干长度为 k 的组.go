/**
@author ZhengHao Lou
Date    2022/01/16
*/
package main

import "fmt"

func divideString(s string, k int, fill byte) []string {
	n := len(s)
	var ans = []string{}
	var i int
	for i = 0; i + k <= n; i += k {
		ans = append(ans, s[i:i+k])
	}
	if n%k != 0 {
		bytes := []byte(s[i:])
		for j := n%k; j < k; j++ {
			bytes = append(bytes, fill)
		}
		ans = append(ans, string(bytes))
	}
	return ans
}

func main() {
	fmt.Println(divideString("abcdefghi", 3, 'x'))
	fmt.Println(divideString("abcdefghij", 3, 'x'))
}
