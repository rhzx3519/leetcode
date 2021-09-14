package main

import (
	"fmt"
	"math"
	"sort"
)

/**
https://leetcode-cn.com/problems/minimum-space-wasted-from-packaging/
 */
var mod = int(math.Pow10(9)) + 7

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	//m := len(boxes)
	n := len(packages)
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + packages[i-1]
	}

	var ans = int64(math.MaxInt64)
	for _, box := range boxes {
		sort.Ints(box)
		var waste int64
		var last int

		if packages[n-1] > box[len(box)-1] {
			continue
		}

		for _, b := range box {
			if last == n {
				break
			}
			i := upperBound(packages, last, b)
			if i == 0 {
				continue
			}
			waste += int64((i - last)*b - (sum[i] - sum[last]))
			last = i
		}

		//fmt.Println(waste)
		if last == n {
			ans = min(ans, waste)
		}
	}

	if ans == int64(math.MaxInt64) {
		return -1
	}
	return int(ans % int64(mod))
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func upperBound(a []int, l, k int) int {
	r := len(a)
	var mid int
	for l < r {
		mid = l + (r-l)>>1
		if a[mid] > k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	fmt.Println(minWastedSpace([]int{2,3,5}, [][]int{{4,8},{2,8}}))
	fmt.Println(minWastedSpace([]int{2,3,5}, [][]int{{1,4},{2,3},{3,4}}))
	fmt.Println(minWastedSpace([]int{3,5,8,10,11,12}, [][]int{{12},{11,9},{10,5,14}}))
}
