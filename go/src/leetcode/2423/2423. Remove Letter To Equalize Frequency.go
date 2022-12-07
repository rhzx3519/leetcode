/**
@author ZhengHao Lou
Date    2022/10/04
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/remove-letter-to-equalize-frequency/
*/
func equalFrequency(word string) bool {
	counter := make([]int, 26)
	for i := range word {
		counter[int(word[i]-'a')]++
	}
	for i := range counter {
		if counter[i] == 0 {
			continue
		}

		freq := make(map[int]int)
		counter[i]--
		for i := range counter {
			if counter[i] > 0 {
				freq[counter[i]]++
			}
		}
		counter[i]++
		if len(freq) == 1 {
			return true
		}
	}
	return false
}

func main() {
	//fmt.Println(equalFrequency("ddaccb"))
	// true
	fmt.Println(equalFrequency("aazz"))
	//
	//fmt.Println(equalFrequency("cbccca"))
}
