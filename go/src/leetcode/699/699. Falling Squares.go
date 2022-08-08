/**
@author ZhengHao Lou
Date    2022/05/26
*/
package main

import "fmt"

/**
https://leetcode.cn/problems/falling-squares/
*/
func fallingSquares(positions [][]int) []int {
	var ans []int
	var mh int
	var intervals [][]int
	for i := range positions {
		l, r, h := positions[i][0], positions[i][0]+positions[i][1], 0
		j := lowerBound(intervals, l+1)
		for j < len(intervals) {
			if intervals[j][0] > r-1 {
				break
			}
			h = max(h, intervals[j][2])
			if intervals[j][0] < l && intervals[j][1] > r {
				rr := intervals[j][1]
				intervals[j][1] = l
				intervals = append(intervals, []int{})
				copy(intervals[j+2:], intervals[j+1:])
				intervals[j+1] = []int{r, rr, intervals[j][2]}
				j++
				break
			} else if intervals[j][0] < l {
				intervals[j][1] = l
				j++
			} else if intervals[j][1] > r {
				intervals[j][0] = r
			} else {
				l = min(l, intervals[j][0])
				r = max(r, intervals[j][1])
				copy(intervals[j:], intervals[j+1:])
				intervals = intervals[:len(intervals)-1]
			}
		}

		intervals = append(intervals, []int{})
		copy(intervals[j+1:], intervals[j:])

		h += positions[i][1]
		intervals[j] = []int{l, r, h}

		mh = max(mh, h)
		ans = append(ans, mh)
	}

	fmt.Println(intervals, ans)
	return ans
}

func lowerBound(intervals [][]int, x int) int {
	l, r := 0, len(intervals)
	for l < r {
		m := l + (r-l)>>1
		if intervals[m][1] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fallingSquares([][]int{{1, 2}, {2, 3}, {6, 1}})
	fallingSquares([][]int{{100, 100}, {200, 100}})
	fallingSquares([][]int{{6, 1}, {9, 2}, {2, 4}}) // [1 2 4]
	fallingSquares([][]int{{9, 7}, {1, 9}, {3, 1}}) // [7,16,17]

	fallingSquares([][]int{{33, 34}, {47, 62}, {70, 16}, {90, 79}, {73, 86}, {55, 6},
		{74, 2}, {40, 95}, {52, 16}, {50, 33}}) // [34,96,112,175,261,261,263,358,374,407]

	fallingSquares([][]int{{1, 6}, {9, 2}, {1, 7}, {3, 2}, {3, 5}}) // [6,6,13,15,20]
}
