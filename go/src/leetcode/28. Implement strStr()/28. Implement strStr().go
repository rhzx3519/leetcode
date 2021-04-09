package main

func strStr(haystack string, needle string) int {
	if needle == haystack {
		return 0
	}
	m, n := len(haystack), len(needle)
	if n > m {
		return -1
	}
	for i := 0; i <= m-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}
