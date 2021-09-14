package main

/**
思路：遍历triplets，取所有元素小于等于target的triplet，做max操作
 */
func mergeTriplets(ts [][]int, ta []int) bool {
	a := []int{0,0,0}
	for _, t := range ts {
		if t[0] <= ta[0] && t[1] <= ta[1] && t[2] <= ta[2] {
			a[0] = max(a[0], t[0])
			a[1] = max(a[1], t[1])
			a[2] = max(a[2], t[2])
		}
	}
	if (a[0] == ta[0] && a[1] == ta[1] && a[2] == ta[2]) {
		return true
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
}

