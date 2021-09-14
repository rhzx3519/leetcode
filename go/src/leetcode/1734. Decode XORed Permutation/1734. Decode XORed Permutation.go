package main

import (
	"fmt"
)
/**
https://leetcode-cn.com/problems/decode-xored-permutation/
思路：perm := []int{a,b,c,d,e}
encoded := []int{a^b,b^c,c^d,d^e}
a^b^c^d^e = t
b^c = encoded[1]
d^c = encoded[3]
a = t ^ encoded[1] ^ encoded[3]
 */
func decode(encoded []int) []int {
	n := len(encoded) + 1
	t := 0
	for i := 1; i <= n; i++  {
		t ^= i
	}

	x := 0
	for i := 1; i < n-1; i += 2 {
		x ^= encoded[i]
	}

	a := t ^ x
	perm := make([]int, n)
	perm[0] = a
	for i := 1; i < n; i++ {
		perm[i] = encoded[i-1]^perm[i-1]
	}
	return perm
}

func main() {
	fmt.Println(decode([]int{3,1}))
	fmt.Println(decode([]int{6,5,4,6}))
}
