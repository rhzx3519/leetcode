/**
@author ZhengHao Lou
Date    2022/05/16
*/
package main

/**
https://leetcode.cn/problems/find-resultant-array-after-removing-anagrams/v
*/
const N = 26

func removeAnagrams(words []string) []string {
	var ws [][]int
	for i := range words {
		var tmp = make([]int, N)
		for _, c := range words[i] {
			tmp[int(c-'a')]++
		}
		ws = append(ws, tmp)
	}
	n := len(words)

	var arr = []int{0}
	for i := 1; i < n; i++ {
		if !is(ws[arr[len(arr)-1]], ws[i]) {
			arr = append(arr, i)
		}
	}
	var ans []string
	for i := range arr {
		ans = append(ans, words[arr[i]])
	}
	//fmt.Println(ans)
	return ans
}

func is(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	//removeAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"})
	//removeAnagrams([]string{"a", "b", "c", "d", "e"})
	// ["meh","iqfgmpec","jawtcmf","fdkbqb"]
	removeAnagrams([]string{"meh", "iqfgmpec", "qefigmpc", "jawtcmf", "fdkbqb"})
}
