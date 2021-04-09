package main

import "fmt"

func trap(height []int) int {
	ans := 0

	l, r := 0, len(height)-1
	for l < r {
		minHeight := min(height[l], height[r])
		if height[l] == minHeight {
			l++
			for l < r && height[l] < minHeight {
				ans += minHeight - height[l]
				l++
			}
		} else {
			r--
			for l < r && minHeight > height[r] {
				ans += minHeight - height[r]
				r--
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
