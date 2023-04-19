package main

import (
	"fmt"
	"strconv"
)

func numDupDigitsAtMostN(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	f := make([][1 << 10]int, m)
	for i := range f {
		for j := range f[i] {
			f[i][j] = -1
		}
	}

	var dfs func(i, mask int, isLimit, isNum bool) int
	dfs = func(i, mask int, isLimit, isNum bool) (tot int) {
		if i == m {
			if isNum {
				return 1
			}
			return
		}
		if !isLimit && isNum {
			x := &f[i][mask]
			if *x >= 0 {
				return *x
			}
			defer func() { *x = tot }()
		}
		if !isNum {
			tot += dfs(i+1, mask, false, false)
		}
		var d int
		if !isNum {
			d = 1
		}
		k := 9
		if isLimit {
			k = int(s[i] - '0')
		}
		for ; d <= k; d++ {
			if mask>>d&1 == 0 {
				tot += dfs(i+1, mask|1<<d, isLimit && d == k, true)
			}
		}
		return
	}

	return n - dfs(0, 0, true, false)
}

func main() {
	fmt.Println(numDupDigitsAtMostN(20))
	fmt.Println(numDupDigitsAtMostN(1000))
}
