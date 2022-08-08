/**
@author ZhengHao Lou
Date    2021/11/28
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
 */
const N = 26
func findAnagrams(s string, p string) []int {
	ns, np := len(s), len(p)
	ms, mp := [N]int{}, [N]int{}
	for i := range p {
		mp[int(p[i] - 'a')]++
	}

	var ans = []int{}
	for i := 0; i < ns; i++ {
		ms[int(s[i] - 'a')]++
		if i < np - 1 {
			continue
		}
		if isEqual(ms, mp) {
			ans = append(ans, i - np + 1)
		}
		ms[int(s[i - np + 1] - 'a')]--
	}

	return ans
}

func isEqual(ms, mp [N]int) bool {
	for i := range ms {
		if ms[i] != mp[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}
