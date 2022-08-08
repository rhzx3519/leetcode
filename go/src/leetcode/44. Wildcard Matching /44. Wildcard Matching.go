package main

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	var mem = make(map[[2]int]bool)

	var f func(int, int) bool
	f = func(i int, j int) bool {
		if _, ok := mem[[2]int{i,j}]; ok {
			return mem[[2]int{i,j}]
		}
		if j == n {
			return i == m
		}

		var b bool
		first := i < m && (p[j] == '?' || p[j] == s[i])
		if p[j] == '*' {
			b = (i < m-1 && f(i+1, j)) || f(i, j+1) || f(i+1, j+1)
		} else {
			b = first && f(i+1, j+1)
		}
		mem[[2]int{i,j}] = b
		return b
	}

	return f(0, 0)
}