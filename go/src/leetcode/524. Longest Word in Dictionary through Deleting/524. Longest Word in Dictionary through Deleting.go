package main

func findLongestWord(s string, dictionary []string) string {
	var maxLen int
	var ans string
	for _, p := range dictionary {
		if isSub(s, p) {
			if len(p) > maxLen {
				maxLen = len(p)
				ans = p
			} else if len(p) == maxLen {
				if p < ans {
					ans = p
				}
			}
		}
	}
	return ans
}

func isSub(s, p string) bool {
	if len(s) < len(p) {
		return false
	}
	var i, j int
	for i < len(s) && j < len(p) {
		if s[i] == p[j] {
			j++
		}
		i++
	}
	return j == len(p)
}
