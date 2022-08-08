/**
@author ZhengHao Lou
Date    2021/10/05
*/
package main

type PeekingIterator struct {
	i int
	arr []int
}

func Constructor(iter *Iterator) *PeekingIterator {
	pi := &PeekingIterator{
		arr: []int{},
	}
	for iter.hasNext() {
		pi.arr = append(pi.arr, iter.next())
	}
	return pi
}

func (this *PeekingIterator) hasNext() bool {
	return this.i < len(this.arr)
}

func (this *PeekingIterator) next() int {
	t := this.arr[this.i]
	this.i++
	return t
}

func (this *PeekingIterator) peek() int {
	return this.arr[this.i]
}

type Iterator struct {

}

func (this *Iterator) hasNext() bool {
	return true
}

func (this *Iterator) next() int {
	return 0
}
