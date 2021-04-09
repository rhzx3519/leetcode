package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow(10, 9)) + 7

func countHomogenous(s string) int {
	count := 0
	last := 0
	for i, _ := range s {
		if i > 0 && s[i] == s[i-1] {
			last++
		} else {
			last = 1
		}
		count += last
		count %= MOD
	}
	return count
}

func main() {
	fmt.Println(countHomogenous("abbcccaa"))
	fmt.Println(countHomogenous("xy"))
	fmt.Println(countHomogenous("zzzzz"))
}
