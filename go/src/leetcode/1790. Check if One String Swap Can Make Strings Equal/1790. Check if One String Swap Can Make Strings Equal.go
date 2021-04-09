package main

func areAlmostEqual(s1 string, s2 string) bool {
	n := len(s1)
	c := []int{}
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			c = append(c, i)
		}
		if len(c) > 2 {
			return false
		}
	}

	if len(c) == 0 {
		return true
	}
	if len(c) != 2 {
		return false
	}
	return s1[c[0]] == s2[c[1]] && s1[c[1]] == s2[c[0]]
}
