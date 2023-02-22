package main

func distinctIntegers(n int) int {
	var cnt = map[int]int{n: 1}
	var last = 1
	for {
		for x := range cnt {
			for i := 1; i <= n; i++ {
				if x%i == 1 {
					cnt[i]++
				}
			}
		}
		if len(cnt) == last {
			break
		}
		last = len(cnt)
	}
	return len(cnt)
}
