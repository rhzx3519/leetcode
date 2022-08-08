/**
@author ZhengHao Lou
Date    2021/10/15
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-and-say/
 */
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	last := countAndSay(n-1)
	bytes := []byte{}
	var ch byte
	var count int
	for i := range last {
		if last[i] != ch {
			if count > 0 {
				bytes = append(bytes, byte('0' + count))
				bytes = append(bytes, ch)
			}
			ch = last[i]
			count = 1
		} else {
			count++
		}
	}
	if count > 0 {
		bytes = append(bytes, byte('0' + count))
		bytes = append(bytes, ch)
	}

	return string(bytes)
}

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(countAndSay(i))
	}
}