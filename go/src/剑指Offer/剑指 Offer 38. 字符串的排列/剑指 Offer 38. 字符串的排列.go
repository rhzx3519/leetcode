package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {
	n := len(s)
	ans := []string{}
	ss := []byte(s)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] < ss[j]
	})

	mapper := make(map[byte]int)
	for i := range ss {
		mapper[ss[i]]++
	}

	var backtrace func([]byte, map[byte]int)
	backtrace = func(bytes []byte, mapper map[byte]int) {
		if len(bytes) == n {
			ans = append(ans, string(bytes))
			return
		}
		for i := range ss {
			if t, ok := mapper[ss[i]]; ok && t <= 0 {
				continue
			}
			if i > 0 && ss[i] == ss[i-1] {
				continue
			}
			mapper[ss[i]]--
			bytes = append(bytes, ss[i])
			backtrace(bytes, mapper)
			bytes = bytes[:len(bytes)-1]
			mapper[ss[i]]++
		}

	}

	backtrace([]byte{}, mapper)
	return ans
}

func main() {
	fmt.Println(permutation("aabc"))
}