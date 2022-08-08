/**
@author ZhengHao Lou
Date    2022/01/05
*/
package main

/**
https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
 */
const N int = 26
func modifyString(s string) string {
	bytes := []byte{}
	for i := range s {
		if s[i] != '?' {
			bytes = append(bytes, s[i])
		} else {
			letters := make([]int, N)
			if i > 0 {
				letters[bytes[i-1] - 'a'] = 1
			}

			if i < len(s) - 1 && s[i+1] != '?' {
				letters[s[i+1] - 'a'] = 1
			}
			for j := range letters {
				if letters[j] == 0 {
					bytes = append(bytes, byte('a' + j))
					break
				}
			}
		}
	}
	return string(bytes)
}

func main() {
	modifyString("j?qg??b")
}
