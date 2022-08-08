/**
@author ZhengHao Lou
Date    2022/08/02
*/
package main

/**
https://leetcode.cn/problems/design-circular-queue/
*/
type MyCircularQueue struct {
	head, tail int
	capacity   int
	arr        []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		arr:      make([]int, k+1),
		capacity: k + 1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr[this.tail] = value
	this.tail = (this.tail + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.capacity
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[(this.tail-1+this.capacity)%this.capacity]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.capacity == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
