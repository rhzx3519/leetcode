package main

import "fmt"

func makeEqual(words []string) bool {
	letters := make([]int, 26)
	nw := len(words)
	for _, w := range words {
		for _, c := range w {
			letters[c-'a']++
		}
	}
	for _, cnt := range letters {
		if cnt % nw != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(makeEqual([]string{"abc","aabc","bc"}))
	fmt.Println(makeEqual([]string{"ab","a"}))
}
