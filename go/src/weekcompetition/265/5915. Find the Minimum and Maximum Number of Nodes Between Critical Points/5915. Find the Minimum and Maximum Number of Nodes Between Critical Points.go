/**
@author ZhengHao Lou
Date    2021/11/01
*/
package main

import (
	"fmt"
	"math"
)

/**
https://leetcode-cn.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
 */
type ListNode struct {
	Val int
    Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	var extrem = []int{}
	var last int
	for p, i := head, 0; p != nil ; p, i = p.Next, i + 1 {
		if 	last != 0 && p.Next != nil {
			if (p.Val > last && p.Val > p.Next.Val) || (p.Val < last && p.Val < p.Next.Val) {
				extrem = append(extrem, i)
			}
		}
		last = p.Val
	}
	if len(extrem) <= 1 {
		return []int{-1, -1}
	}
	var minLen = math.MaxInt32 >> 1
	for i := 0; i < len(extrem) - 1; i++ {
		minLen = min(minLen, extrem[i + 1] - extrem[i])
	}
	ans := []int{minLen, extrem[len(extrem) - 1] - extrem[0]}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(nodesBetweenCriticalPoints(nil))
}

