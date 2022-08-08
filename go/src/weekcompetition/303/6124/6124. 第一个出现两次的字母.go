/**
@author ZhengHao Lou
Date    2022/07/24
*/
package main

func repeatedCharacter(s string) byte {
	counter := make(map[byte]int)
	for i := range s {
		counter[s[i]]++
		if counter[s[i]] == 2 {
			return s[i]
		}
	}
	return 'a'
}
