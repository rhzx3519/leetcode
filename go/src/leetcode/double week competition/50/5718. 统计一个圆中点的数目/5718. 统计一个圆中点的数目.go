package main

import "fmt"

func countPoints(points [][]int, queries [][]int) []int {
	ans := make([]int, len(queries))
	
	for i, q := range queries {
		x, y, r := q[0], q[1], q[2]
		ans[i] = in(points, x, y, r*r)
	}
	
	return ans
}

func in(points [][]int, x, y, r2 int) int {
	count := 0

	for _, point := range points {
		i, j := point[0], point[1]
		if (x-i)*(x-i) + (y-j)*(y-j) <= r2 {
			count++
		}
	}

	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(countPoints([][]int{{1,3},{3,3},{5,3},{2,2}}, [][]int{{2,3,1}, {4,3,1}, {1,1,2}}))
	fmt.Println(countPoints([][]int{{1,1},{2,2},{3,3},{4,4},{5,5}}, [][]int{{1,2,2}, {2,2,2}, {4,3,2}, {4,3,3}}))
}