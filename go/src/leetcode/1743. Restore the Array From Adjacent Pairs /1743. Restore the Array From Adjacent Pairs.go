package main

import "fmt"

func restoreArray(adjacentPairs [][]int) []int {
	n := len(adjacentPairs) + 1
	nums := make([]int, n)

	mapper := make(map[int][]int)


	for _, adj := range adjacentPairs {
		a, b := adj[0], adj[1]
		if _, ok := mapper[a]; !ok {
			mapper[a] = []int{}
		}
		if _, ok := mapper[b]; !ok {
			mapper[b] = []int{}
		}
		mapper[a] = append(mapper[a], b)
		mapper[b] = append(mapper[b], a)
	}
	tmp := []int{}
	for num, adj := range mapper {
		if len(adj) == 1 {
			tmp = append(tmp, num)
		}
	}
	nums[0], nums[n-1] = tmp[0], tmp[1]

	for i := 1; i < n-1; i++ {
		for _, t := range mapper[nums[i-1]] {
			if i == 1 || (i > 1 && t != nums[i-2]) {
				nums[i] = t
				break
			}
		}
	}

	return nums
}

func main() {
	fmt.Println(restoreArray([][]int{{2,1},{3,4},{3,2}}))
	fmt.Println(restoreArray([][]int{{4,-2},{1,4},{-3,1}}))
	fmt.Println(restoreArray([][]int{{100000,-100000}}))
}
