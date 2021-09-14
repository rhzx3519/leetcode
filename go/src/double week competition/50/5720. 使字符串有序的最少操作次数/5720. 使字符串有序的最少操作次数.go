package main

import (
	"fmt"
	"math"
)

var MOD = int(math.Pow(10, 9)) + 7

func makeStringSorted(s string) int {
	count := 0
	a := []byte(s)
	n := len(a)
	walk:
	for i := n-1; i > 0; i-- {

		if a[i] < a[i-1] {
			j := i
			for ; j < len(a) && a[j] < a[i-1]; j++ {}
			j--
			a[i-1], a[j] = a[j], a[i-1]
			reverse(a, i)
			count = (count + 1) % MOD
			goto walk
		}
	}

	return count
}


func reverse(a []byte, i int) {
	l, r := i, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
}

func main() {
	//fmt.Println(makeStringSorted("cba"))
	//fmt.Println(makeStringSorted("aabaa"))
	//fmt.Println(makeStringSorted("cdbea"))
	fmt.Println(makeStringSorted("leetcodeleetcodeleetcode"))
}
