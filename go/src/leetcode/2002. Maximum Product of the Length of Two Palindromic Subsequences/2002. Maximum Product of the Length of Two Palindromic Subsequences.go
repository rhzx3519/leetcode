package main

import "fmt"

/**
思路；状态压缩DP
f[mask] 表示状态是mask时，最长的回文子序列的长度
```code
// 求mask状态的所有子状态
for subset := mask; subset !=0; subset = (subset - 1)&mask {}
```

 */
func maxProduct(s string) int {
	n := len(s)
	f := make([]int, 1<<n)
	valid := make([]bool, 1<<n)
	for mask := 1; mask < 1<<n; mask++ {
		sz := palindromic(s, mask)
		if sz == 0 {
			continue
		}
		f[mask] = sz
		valid[mask] = true
	}

	var ans int
	for mask := 1; mask < 1<<n; mask++ {
		if !valid[mask] {
			continue
		}
		tmp := mask ^ (1<<n - 1)
		for subset := tmp; subset != 0; subset = (subset - 1)&tmp {
			if !valid[subset] {
				continue
			}
			ans = max(ans, f[mask] * f[subset])
		}
	}

	return ans
}

// mask 表示s的子序列的状态，如果是回文串，返回回文串的长度
func palindromic(s string, mask int) int {
	var sz int
	l, r := 0, len(s) - 1
	for l <= r {
		for l <= r && mask & (1<<l) == 0 {
			l++
		}
		for l <= r && mask & (1<<r) == 0 {
			r--
		}
		if l > r {
			break
		}
		if s[l] != s[r] {
			return 0
		}
		if l == r {
			sz++
		} else {
			sz += 2
		}
		l++
		r--
	}

	return sz
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProduct("leetcodecom"))
	fmt.Println(maxProduct("bb"))
	fmt.Println(maxProduct("accbcaxxcxx"))
	fmt.Println(maxProduct("zyz"))
}