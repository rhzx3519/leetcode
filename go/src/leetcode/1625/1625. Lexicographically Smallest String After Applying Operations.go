package main

/*
*
https://leetcode.cn/problems/lexicographically-smallest-string-after-applying-operations/description/
*/
func findLexSmallestString(s string, a int, b int) string {
	n := len(s)
	ans := s
	for i := 0; i < n; i++ { // 最多轮转n次会重复
		s = s[n-b:] + s[:n-b]
		cs := []byte(s)
		for j := 0; j < 10; j++ {
			for k := 1; k < n; k += 2 {
				cs[k] = byte('0' + (int(cs[k]-'0')+a)%10)
			}
			if b&1 == 1 {
				for p := 0; p < 10; p++ {
					for k := 0; k < n; k += 2 {
						cs[k] = byte('0' + (int(cs[k]-'0')+a)%10)
					}
					s = string(cs)
					if ans > s {
						ans = s
					}
				}
			} else {
				s = string(cs)
				if s < ans {
					ans = s
				}
			}
		}
	}
	return ans
}
