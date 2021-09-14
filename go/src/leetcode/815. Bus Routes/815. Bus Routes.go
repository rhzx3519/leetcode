package main

import "fmt"

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}
	n := len(routes)
	visited := make([]bool, n)
	que := []int{}

	stage2bus := make(map[int][]int)


	for b, route := range routes {
		for _, stage := range route {
			if _, ok := stage2bus[stage]; !ok {
				stage2bus[stage] = []int{}
			}
			stage2bus[stage] = append(stage2bus[stage], b)
		}
	}

	for _, b := range stage2bus[source] {
		visited[b] = true
		que = append(que, b)
	}

	var step int
	for len(que) > 0 {
		step++
		for sz := len(que); sz > 0; sz-- {
			i := que[0]
			que = que[1:]

			for _, ni := range routes[i] {
				if ni == target {
					return step
				}
				for _, b := range stage2bus[ni] {
					if visited[b] {
						continue
					}
					visited[b] = true
					que = append(que, b)
				}

			}
		}
	}

	return -1
}

func main() {
	fmt.Println(numBusesToDestination([][]int{{1,2,7},{3,6,7}}, 1, 6))
	fmt.Println(numBusesToDestination([][]int{{7,12},{4,5,15},{6},{15,19},{9,12,13}}, 15, 12))
}