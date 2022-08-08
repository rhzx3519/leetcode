/**
@author ZhengHao Lou
Date    2022/06/26
*/
package main

import "math/rand"

/**
https://leetcode.cn/problems/random-pick-with-blacklist/
*/
type Solution struct {
	b2w   map[int]int
	bound int
}

func Constructor(n int, blacklist []int) Solution {
	m := len(blacklist)
	bound := n - m

	black := make(map[int]bool)
	for _, i := range blacklist {
		if i >= bound {
			black[i] = true
		}
	}

	b2w := make(map[int]int)
	w := bound
	for _, i := range blacklist {
		if i < bound {
			for black[w] {
				w++
			}
			b2w[i] = w
			w++
		}
	}
	return Solution{b2w: b2w, bound: bound}
}

func (this *Solution) Pick() int {
	x := rand.Intn(this.bound)
	if this.b2w[x] != 0 {
		return this.b2w[x]
	}
	return x
}
