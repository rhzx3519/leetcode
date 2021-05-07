package main

func minInterval(intervals [][]int, queries []int) []int {
	ans := []int{}
	qs := make([][]int, len(queries))
	for i, q := range queries {
		qs[i] = []int{i, q}
	}
	return ans
}
