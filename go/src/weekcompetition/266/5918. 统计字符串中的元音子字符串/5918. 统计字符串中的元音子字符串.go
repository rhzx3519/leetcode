/**
@author ZhengHao Lou
Date    2021/11/07
*/
package main

import "fmt"

func countVowelSubstrings(word string) int {
	n := len(word)
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var x int
		if vowels[word[i-1]] !=  0 {
			x = 1
		}
		sum[i] = sum[i-1] + x
	}

	var ans int64
	for i := 0; i < n; i++ {
		for j := i+1; j <= n; j++ {
			ans += int64(sum[j] - sum[i])
		}
	}

	return int(ans)
}

func main() {
	fmt.Println(countVowelSubstrings("aba"))
	fmt.Println(countVowelSubstrings("abc"))
	fmt.Println(countVowelSubstrings("ltcd"))
	fmt.Println(countVowelSubstrings("noosabasboosa"))
}
