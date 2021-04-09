package main

import "fmt"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	notConnected := make(map[int]int)
	mostLanguage := make(map[int]int)

	for _, f := range friendships {
		a, b := f[0], f[1]
		if same(a, b, languages) {
			continue
		}
		notConnected[a]++
		notConnected[b]++
	}

	for f := range notConnected {
		for _, l := range languages[f-1] {
			mostLanguage[l]++
		}
	}

	// 最多人会的语言，有多少人会
	x := 0
	for _, v := range mostLanguage {
		if v > x {
			x = v
		}
	}

	return len(notConnected) - x
}

func same(a, b int, languages [][]int) bool {
	for _, l := range languages[a-1] {
		for _, r := range languages[b-1] {
			if l == r {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println(minimumTeachings(2, [][]int{{1},{2},{1,2}}, [][]int{{1,2},{1,3},{2,3}}))
}