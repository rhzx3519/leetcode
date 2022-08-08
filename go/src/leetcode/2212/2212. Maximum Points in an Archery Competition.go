/**
@author ZhengHao Lou
Date    2022/03/21
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-points-in-an-archery-competition/
思路：dfs
time: O(2^11)
space: O(2^11)
*/
func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	n := len(aliceArrows)
	var bobArrows []int
	var mx int

	var dfs func(k, left, score int, bob []int)
	dfs = func(k, left, score int, bob []int) {
		if k == n || left == 0 {
			if score > mx {
				mx = score
				bobArrows = append([]int{}, bob...)
				bobArrows[n-1] += left
			}
			return
		}
		if left > aliceArrows[k] {
			bob[k] = aliceArrows[k] + 1
			dfs(k+1, left-aliceArrows[k]-1, score+k, bob)
			bob[k] = 0
		}

		dfs(k+1, left, score, bob)
	}

	dfs(0, numArrows, 0, make([]int, n))

	fmt.Println(mx, bobArrows)
	return bobArrows
}

func main() {
	//maximumBobPoints(9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2, 0})
	//maximumBobPoints(3, []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2})
	maximumBobPoints(1, []int{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0})
}
