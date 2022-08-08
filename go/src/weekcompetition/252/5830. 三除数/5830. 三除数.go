package main

func isThree(n int) bool {
	var cnt int
	for i := 2; i < n; i++ {
		if n%i == 0 {
			cnt++
		}
		if cnt > 1 {
			return false
		}
	}
	return cnt==1
}
