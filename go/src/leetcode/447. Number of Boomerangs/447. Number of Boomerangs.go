package main

import "fmt"

/**
time: O(n^2)
 */
func numberOfBoomerangs(points [][]int) int {
	n := len(points)
	var count int
	var a = make([]map[int][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make(map[int][]int)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			d := dist(points[i], points[j])
			a[i][d] = append(a[i][d], j)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			d := dist(points[i], points[j])
			if t, ok := a[j][d]; ok {
				count += len(t) - 1
			}
		}
	}

	fmt.Println(count)
	return count
}

func dist(p1, p2 []int) int {
	return (p1[0] - p2[0])*(p1[0] - p2[0]) + (p1[1] - p2[1])*(p1[1] - p2[1])
}

func main() {
	numberOfBoomerangs([][]int{{0,0},{1,0},{2,0}})
	numberOfBoomerangs([][]int{{1,1},{2,2},{3,3}})
	numberOfBoomerangs([][]int{{1,1}})
}