package main

import (
	"fmt"
	"math"
)

func beautySum(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		count := make(map[uint8]int)
		for j := i; j < n; j++ {
			count[s[j]]++
			minVal, maxVal := get(count)
			ans += maxVal - minVal
		}
	}
	return ans
}

func get(count map[uint8]int) (int, int) {
	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, num := range count {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}
	return minVal, maxVal
}

/*
func beautySum(s string) int {
	//count := make(map[rune]int)
	last := 0
	groups := []string{}
	for _, c := range s  {
		tmp := []string{}
		for i := last; i < len(groups); i++ {
			tmp = append(tmp, groups[i] + string(c))
		}
		tmp = append(tmp, string(c))
		last = len(groups)
		groups = append(groups, tmp...)
	}
	fmt.Println(groups)

	ans := 0
	for i := 0; i < len(groups); i++ {
		str := groups[i]
		ans += beautyVal(str)

	}
	return ans
}
 */

func beautyVal(s string) int {
	count := make(map[rune]int)
	minVal, maxVal := math.MaxInt32, math.MinInt32
	for _, c := range s {
		count[c]++;
	}
	if len(count) <= 1 {
		return 0
	}
	for _, num := range count {
		minVal = min(minVal, num)
		maxVal = max(maxVal, num)
	}
	return maxVal - minVal
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func main() {
	fmt.Println(beautySum("aabcb"))
	fmt.Println(beautySum("aabcbaa"))
}