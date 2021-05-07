package main
/**
思路：统计每一层缝隙的位置，最少穿过的砖等于层数-缝隙最多的数量
 */
func leastBricks(wall [][]int) int {
	mapper := make(map[int]int)
	for _, w := range wall {
		s := 0
		for i := 0; i < len(w)-1; i++ {
			s += w[i]
			mapper[s]++
		}
	}

	maxVal := 0
	for _, d := range mapper {
		if d > maxVal {
			maxVal = d
		}
	}

	return len(wall) - maxVal
}

func main() {
	leastBricks([][]int{{1,2,2,1},{3,1,2},{1,3,2},{2,4},{3,1,2},{1,3,1,1}})
}