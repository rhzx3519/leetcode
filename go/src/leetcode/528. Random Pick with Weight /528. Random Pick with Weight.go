package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	ws []int
	s int
}


func Constructor(w []int) Solution {
	ws := []int{}
	var s int
	for _, t := range w {
		s += t
		ws = append(ws, s)
	}
	return Solution{
		ws: ws,
		s: s,
	}
}


func (this *Solution) PickIndex() int {
	target := rand.Int31n(int32(this.s)) + 1
	i := lowerBound(this.ws, int(target))
	fmt.Println(i)
	return i
}

func lowerBound(ws []int, target int) int {
	l, r := 0, len(ws)
	var m int
	for l < r {
		m = l + (r-l)>>1
		if ws[m] >= target {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	obj := Constructor([]int{1,3})
	obj.PickIndex()

}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
