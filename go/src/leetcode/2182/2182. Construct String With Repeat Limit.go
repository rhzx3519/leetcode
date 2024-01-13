/*
*
@author ZhengHao Lou
Date    2022/02/21
*/
package main

import "fmt"

/*
*
https://leetcode-cn.com/problems/construct-string-with-repeat-limit/
*/
const N = 26

func repeatLimitedString(s string, repeatLimit int) string {
	var cnt = [26]int{}
	for i := range s {
		cnt[int(s[i]-'a')]++
	}
	var ans []byte
OUT:
	for i := N - 1; i >= 0; i-- {
		var j = repeatLimit
		for ; cnt[i] > 0 && j > 0; j-- {
			ans = append(ans, byte('a'+i))
			cnt[i]--
		}
		if cnt[i] == 0 {
			continue
		}
		for k := i - 1; k >= 0; k-- {
			if cnt[k] > 0 {
				ans = append(ans, byte('a'+k))
				cnt[k]--
				goto OUT
			}
		}
		break
	}
	return string(ans)
}

func main() {
	fmt.Println(repeatLimitedString("aababab", 2))
}
