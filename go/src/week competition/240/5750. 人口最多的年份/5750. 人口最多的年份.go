package main

import (
	"fmt"
)

/**
思路：
1。 每次获取一个区间时，将区间内的值都+1，time: O(n^2), space: O(1)
2. 差分数组，只对区间头尾+1/-1, time: O(n), space: O(1)
 */
func maximumPopulation(logs [][]int) int {
	var N = 101
	var base = 1950
	diff := make([]int, N)
	for _, log := range logs {
		diff[log[0] - base]++
		diff[log[1] - base]--
	}
	var count int
	var year, maxP int
	for i, p := range diff {
		count += p
		if count > maxP {
			maxP = count
			year = i
		}
	}
	return year + base
}

func main() {
	fmt.Println(maximumPopulation([][]int{{1993,1999}, {2000,2010}}))
	fmt.Println(maximumPopulation([][]int{{1950,1961}, {1960,1971}, {1970,1981}}))
}


