package main

func areOccurrencesEqual(s string) bool {
	bytes := make([]int, 26)
	for _, c := range s {
		bytes[int(c - 'a')]++
	}
	var k int
	for _, i := range bytes {
		if i == 0 {
			continue
		}
		if k != 0 && k != i {
			return false
		}
		k = i
	}
	return true
}
