/**
@author ZhengHao Lou
Date    2022/09/23
*/
package main

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

type MyLinkedList struct {
	Head, Tail *LinkedNode
	Size       int
}

func Constructor() MyLinkedList {
	node := &LinkedNode{}
	return MyLinkedList{
		Head: node,
		Tail: node,
	}
}

func (this *MyLinkedList) Get(index int) int {
	h := this.Head
	for i := 0; i < index; i++ {
		h = h.Next
	}
	if h == nil || h.Next == nil {
		return -1
	}
	return h.Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &LinkedNode{
		Val: val,
	}
	node.Next = this.Head.Next
	this.Head.Next = node
	if this.Tail == this.Head {
		this.Tail = node
	}
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &LinkedNode{
		Val: val,
	}
	this.Tail.Next = node
	this.Tail = node
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		this.AddAtHead(val)
		return
	}
	if index > this.Size {
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	h := this.Head
	for i := 0; i < index; i++ {
		h = h.Next
	}
	node := &LinkedNode{
		Val: val,
	}
	node.Next = h.Next
	h.Next = node
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}
	h := this.Head
	for i := 0; i < index; i++ {
		h = h.Next
	}
	if h.Next == this.Tail {
		this.Tail = h
		h.Next = nil
		this.Size--
		return
	}
	h.Next = h.Next.Next
	this.Size--
}

func main() {
	linkedList := Constructor()
	linkedList.AddAtHead(1)
	linkedList.AddAtTail(3)
	linkedList.AddAtIndex(1, 2) //链表变为1-> 2-> 3
	linkedList.Get(1)           //返回2
	linkedList.DeleteAtIndex(2) //现在链表是1-> 3
	linkedList.Get(1)           //返回3
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
