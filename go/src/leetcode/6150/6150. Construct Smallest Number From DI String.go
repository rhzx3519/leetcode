/**
@author ZhengHao Lou
Date    2022/08/15
*/
package main

/**
https://leetcode.cn/problems/construct-smallest-number-from-di-string/
*/
func smallestNumber(pattern string) string {
	pattern += "I"
	var bytes []byte
	var st []int
	var c = 1
	for i := range pattern {
		if pattern[i] == 'D' {
			st = append(st, c)
			c++
		} else {
			bytes = append(bytes, byte(c+'0'))
			for len(st) != 0 {
				j := st[len(st)-1]
				st = st[:len(st)-1]
				bytes = append(bytes, byte(j+'0'))
			}
			c++
		}
	}
	
	//fmt.Println(c, string(bytes))
	return string(bytes)
}

func main() {
	smallestNumber("IIIDIDDD")
	smallestNumber("DDD")
}
