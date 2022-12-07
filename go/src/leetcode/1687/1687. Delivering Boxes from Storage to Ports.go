/*
*
@author ZhengHao Lou
Date    2022/12/05
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/delivering-boxes-from-storage-to-ports/description/
f[i]表示运送前i个箱子所需的最少行程次数
f[i] = f[j] + sum(cost[j+1:i])
*/
func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	ws, cs := make([]int, n+1), make([]int, n+1)
	for i, box := range boxes {
		p, w := box[0], box[1]
		ws[i+1] = ws[i] + w
		if i < n-1 {
			var t int
			if p != boxes[i+1][0] {
				t++
			}
			cs[i+1] = cs[i] + t
		}
	}

	f := make([]int, n+1)
	// 1. 遍历求f[j] - cs[j]的最小值
	//for i := 1; i <= n; i++ {
	//	f[i] = 1 << 30
	//	for j := max(0, i-maxBoxes); j < i; j++ {
	//		if ws[i]-ws[j] <= maxWeight {
	//			f[i] = min(f[i], f[j]+cs[i-1]-cs[j]+2)
	//		}
	//	}
	//}
	// 2. 单调队列
	var q = []int{0}
	for i := 1; i <= n; i++ {
		f[i] = 1 << 30
		for len(q) != 0 && (i-q[0] > maxBoxes || ws[i]-ws[q[0]] > maxWeight) {
			q = q[1:]
		}
		f[i] = min(f[i], f[q[0]]-cs[q[0]]+cs[i-1]+2)

		if i < n {
			for len(q) != 0 && f[i]-cs[i] <= f[q[len(q)-1]]-cs[q[len(q)-1]] {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
	}

	fmt.Println(ws, cs, f[n])
	return f[n]
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
	boxDelivering([][]int{{1, 1}, {2, 1}, {1, 1}}, 2, 3, 3)
	boxDelivering([][]int{{1, 2}, {3, 3}, {3, 1}, {3, 1}, {2, 4}}, 3, 3, 6)
}
