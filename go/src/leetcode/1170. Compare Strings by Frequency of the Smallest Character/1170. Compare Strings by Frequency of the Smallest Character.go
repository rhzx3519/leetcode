package main

import (
	"fmt"
	"sort"
)

func numSmallerByFrequency(queries []string, words []string) []int {
	n := len(words)
	ws := []int{}
	for _, w := range words {
		ws = append(ws, f(w))
	}
	sort.Ints(ws)

	ans := []int{}
	for _, q := range queries {
		i := UpperBound(ws, f(q))
		ans = append(ans, n - i)
	}
	return ans
}

func f(s string) int {
	minC := byte('z')
	mapper := make(map[byte]int)
	for _, c := range s {
		mapper[byte(c)]++
		if byte(c) < minC {
			minC = byte(c)
		}
	}
	return mapper[minC]
}

/**
返回第一个arr中第一个大于k的元素下标，
如果k大于arr中的所有元素，则返回arr的长度
*/
func UpperBound(arr []int, k int) int {
	l, r := 0, len(arr)
	var mid int
	for l < r {
		mid = l + (r - l)>>1
		if arr[mid] > k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Println(numSmallerByFrequency([]string{"bbb","cc"}, []string{"a","aa","aaa","aaaa"}))
}
