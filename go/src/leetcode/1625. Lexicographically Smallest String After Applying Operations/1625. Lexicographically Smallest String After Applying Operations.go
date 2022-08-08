/**
@author ZhengHao Lou
Date    2021/11/05
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/lexicographically-smallest-string-after-applying-operations/
思路：暴力BFS
 */
func findLexSmallestString(s string, a int, b int) string {
	var ans = ""
	//n := len(s)
	vis := make(map[string]bool)
	vis[s] = true
	que := []string{s}
	for len(que) > 0 {
		s = que[0]
		que = que[1:]
		if ans == "" || s < ans {
			ans = s
		}
		//fmt.Println(s)
		t := s
		for {
			bs := []byte{}
			for i := range t {
				if i&1 == 1 {
					bs = append(bs, byte('0' + (int(s[i] - '0') + a) % 10))
				} else {
					bs = append(bs, t[i])
				}
			}
			t = string(bs)
			if vis[t] {
				break
			}
			que = append(que, t)
			vis[t] = true
		}

		t = s
		for {
			bs := []byte{}
			bs = append(bs, []byte(t)[b:]...)
			bs = append(bs, []byte(t)[:b]...)
			t = string(bs)
			if vis[t] {
				break
			}
			que = append(que, t)
			vis[t] = true
		}
	}

	return ans
}

func main() {
	fmt.Println(findLexSmallestString("5525", 9, 2))
	fmt.Println(findLexSmallestString("74", 5, 1))
	fmt.Println(findLexSmallestString("0011", 4, 2))
	fmt.Println(findLexSmallestString("43987654", 7, 3))
}