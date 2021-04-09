package main

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	ind := make([]int, n+1)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		ind[a]++
		ind[b]++
		if ind[a] == n-1 {
			return a
		}
		if ind[b] == n-1 {
			return b
		}
	}
	return -1
}
