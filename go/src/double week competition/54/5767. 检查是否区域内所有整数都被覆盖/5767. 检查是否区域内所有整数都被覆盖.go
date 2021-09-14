package main

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		var f bool
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				f = true
				break
			}
		}
		if !f {
			return false
		}
	}
	return true
}