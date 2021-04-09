package main

type MyQueue struct {
	a, b []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.a = append(this.a, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.b) == 0 {
		for len(this.a) > 0 {
			this.b = append(this.b, this.a[len(this.a)-1])
			this.a = this.a[:len(this.a)-1]
		}
	}

	if len(this.b) == 0 {
		return -1
	}
	ans := this.b[len(this.b) - 1]
	this.b = this.b[:len(this.b)-1]
	return ans
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.b) == 0 {
		for len(this.a) > 0 {
			this.b = append(this.b, this.a[len(this.a)-1])
			this.a = this.a[:len(this.a)-1]
		}
	}
	if len(this.b) == 0 {
		return -1
	}
	return this.b[len(this.b) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.a)==0 && len(this.b)==0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
