package main

import (
	"fmt"
	"math"
)

func maxPoints(points [][]int) int {
	var ans int
	n := len(points)
	for i := 0; i < n; i++ {
		mapper := make(map[float32]int)
		for j := i+1; j < n; j++ {
			mapper[slope(points[i], points[j])]++
		}
		var maxVal int
		for _, v := range mapper {
			maxVal = max(maxVal, v)
		}
		ans = max(ans, maxVal+1)
	}

	fmt.Println(ans)
	return ans
}

func slope(a, b []int) float32 {
	var x = a[0] - b[0]
	var y = a[1] - b[1]
	if x == 0 {
		return math.MaxFloat32
	}
	return float32(y) / float32(x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func main() {
	maxPoints([][]int{{1,1},{2,2},{3,3}})
	maxPoints([][]int{{1,1},{3,2},{5,3},{4,1},{2,3},{1,4}})
}

