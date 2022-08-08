/**
@author ZhengHao Lou
Date    2022/03/11
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/count-nodes-with-the-highest-score/
*/
func countHighestScoreNodes(parents []int) int {
	n := len(parents)

	tree := make([][]int, n)
	counter := make([]int, n)
	out := make([]int, n)
	for i := range parents {
		if parents[i] != -1 {
			out[parents[i]]++
			tree[parents[i]] = append(tree[parents[i]], i)
		}
	}

	var que []int
	for i := range out {
		if out[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) != 0 {
		x := que[0]
		que = que[1:]
		p := parents[x]
		if x == 0 {
			continue
		}
		counter[p] += counter[x] + 1
		if out[p]--; out[p] == 0 {
			que = append(que, p)
		}
	}

	var score, c int
	for i := range parents {
		if counter[0]-counter[i] == 0 && len(tree[i]) == 0 {
			continue
		}
		var s = 1
		if counter[0]-counter[i] != 0 {
			s *= counter[0] - counter[i]
		}
		if len(tree[i]) != 0 {
			for _, j := range tree[i] {
				s *= counter[j] + 1
			}
		}

		if s > score {
			score = s
			c = 1
		} else if s == score {
			c++
		}
	}

	fmt.Println(out)
	fmt.Println(counter)
	fmt.Println(c)

	return c
}

func main() {
	countHighestScoreNodes([]int{-1, 2, 0, 2, 0})
	countHighestScoreNodes([]int{-1, 2, 0})
}
