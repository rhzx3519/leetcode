/**
@author ZhengHao Lou
Date    2022/08/15
*/
package main

type MyCircularDeque struct {
	front, end int
	k          int
	que        []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		que: make([]int, k+1),
		k:   k + 1,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front - 1 + this.k) % this.k
	this.que[this.front] = value
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.que[this.end] = value
	this.end = (this.end + 1) % this.k
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % this.k
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.end = (this.end - 1 + this.k) % this.k
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.que[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	i := (this.end - 1 + this.k) % this.k
	return this.que[i]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.front == this.end
}

func (this *MyCircularDeque) IsFull() bool {
	return (this.end+1)%this.k == this.front
}

func main() {
	que := Constructor(4)
	que.InsertLast(9)
	que.DeleteLast()
	que.GetRear()
	que.GetFront()
	
}
