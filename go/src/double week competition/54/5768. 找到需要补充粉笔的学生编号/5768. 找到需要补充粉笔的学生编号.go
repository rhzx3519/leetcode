package main

func chalkReplacer(chalk []int, k int) int {
	var s int
	for _, c := range chalk {
		s += c
	}
	k %= s
	for i, c := range chalk {
		if k < c {
			return i
		}
		k -= c
	}
	return -1
}
