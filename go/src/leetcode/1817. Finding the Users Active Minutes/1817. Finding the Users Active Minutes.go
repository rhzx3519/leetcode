package main

import "fmt"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	ans := make([]int, k)

	mapper := make(map[int]map[int]bool)
	for	_, log := range logs {
		id, t := log[0], log[1]
		if _, ok := mapper[id]; !ok {
			mapper[id] = make(map[int]bool)
		}
		mapper[id][t] = true
	}

	for _, t := range mapper {
		if len(t)-1 < k {
			ans[len(t)-1]++
		}
	}

	return ans
}

func main() {
	fmt.Println(findingUsersActiveMinutes([][]int{{1,1},{2,2},{2,3}}, 4))
}
