/**
@author ZhengHao Lou
Date    2021/11/07
*/
package main

import "fmt"

func countVowels(word string) int64 {
	var ans int64
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

	fmt.Println(sum)
	for i := 0; i < n; i++ {
		if vowels[word[i]] !=  0 {
			ans += int64(n)
		}
	}

	return ans
}

func main() {
	fmt.Println(countVowels("aba"))
	fmt.Println(countVowels("abc"))
	//fmt.Println(countVowels("ltcd"))
	//fmt.Println(countVowels("noosabasboosa"))
}

