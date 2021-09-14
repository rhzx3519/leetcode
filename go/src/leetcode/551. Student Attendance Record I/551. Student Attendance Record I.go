package main

func checkRecord(s string) bool {
	var a, l, maxL int
	for i := range s {
		switch s[i] {
		case 'A':
			a++
			l = 0
		case 'L':
			l++
		default:
			l = 0
		}
		maxL = max(maxL, l)
	}
	return a < 2 && maxL < 3
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
