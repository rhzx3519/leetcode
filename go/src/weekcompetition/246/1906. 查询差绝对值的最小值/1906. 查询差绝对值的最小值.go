package main

import "fmt"

const N = 101

func minDifference(nums []int, queries [][]int) []int {
	n := len(nums)
	pre := make([][]int, n+1)
	pre[0] = make([]int, N)

	for i := 1; i <= n; i++ {
		cnt := make([]int, N)
		copy(cnt, pre[i-1])
		cnt[nums[i-1]]++
		pre[i] = cnt
	}

	ans := []int{}
	for _, q := range queries {
		l, r := q[0], q[1]
		ans = append(ans, absmin(pre, l, r))
	}

	fmt.Println(ans)
	return ans
}

func absmin(pre [][]int, l, r int) int {
	var diff = N
	var last = -1
	for i := 1; i < N; i++ {
		t := pre[r+1][i] - pre[l][i]
		if t > 0 {
			if last != -1 {
				diff = min(diff, i - last)
			}
			last = i
		}
	}
	if diff == N {
		return -1
	}
	return diff
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minDifference([]int{1,3,4,8}, [][]int{{0,1},{1,2},{2,3},{0,3}})
}
