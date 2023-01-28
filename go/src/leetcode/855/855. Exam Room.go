/*
*
@author ZhengHao Lou
Date    2022/12/30
*/
package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/exam-room/
*/
type ExamRoom struct {
	n     int
	seats []int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{
		n:     n,
		seats: []int{},
	}
}

func (this *ExamRoom) Seat() int {
	if len(this.seats) == 0 {
		this.seats = append(this.seats, 0)
		return 0
	}
	var j, v int
	var d = this.seats[0]
	for i := 1; i < len(this.seats); i++ {
		if (this.seats[i]-this.seats[i-1])/2 > d {
			d = (this.seats[i] - this.seats[i-1]) / 2
			j = i
			v = this.seats[i-1] + d
		}
	}
	if this.n-1-this.seats[len(this.seats)-1] > d {
		j = len(this.seats)
		v = this.n - 1
	}

	this.seats = append(this.seats, 0)
	copy(this.seats[j+1:], this.seats[j:])
	this.seats[j] = v
	fmt.Println(j, this.seats)
	return v
}

func (this *ExamRoom) Leave(p int) {
	j := sort.Search(len(this.seats), func(i int) bool {
		return this.seats[i] >= p
	})
	fmt.Println(this.seats[j])
	copy(this.seats[j:], this.seats[j+1:])
	this.seats = this.seats[:len(this.seats)-1]
	fmt.Println(this.seats)
}

func main() {
	// null,0,9,4,null,null,0,4,2,6,1,3,5,7,8,null,null,0,4,
	obj := Constructor(10)
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Leave(0)
	obj.Leave(4)
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Seat()
	obj.Leave(0)
	obj.Leave(4)
	obj.Seat()
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */
