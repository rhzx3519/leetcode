/*
*
@author ZhengHao Lou
Date    2022/12/12
*/
package main

import "fmt"

/*
*
https://leetcode.cn/problems/design-memory-allocator/
*/
type blob struct {
	mID   int
	i, sz int
}

type Allocator struct {
	blobs []blob
	n     int
}

func Constructor(n int) Allocator {
	return Allocator{
		n:     n,
		blobs: []blob{{-1, 0, 0}, {-1, n, 0}},
	}
}

func (this *Allocator) Allocate(size int, mID int) int {
	var free int
	for i := 0; i < len(this.blobs)-1; i++ {
		free = this.blobs[i+1].i - (this.blobs[i].i + this.blobs[i].sz)
		if free >= size {
			this.blobs = append(this.blobs, blob{})
			copy(this.blobs[i+2:], this.blobs[i+1:])
			this.blobs[i+1] = blob{mID: mID, i: this.blobs[i].i + this.blobs[i].sz, sz: size}
			return this.blobs[i+1].i
		}
	}

	return -1
}

func (this *Allocator) Free(mID int) (sz int) {
	for i := 0; i < len(this.blobs); {
		if mID == this.blobs[i].mID {
			sz += this.blobs[i].sz
			copy(this.blobs[i:], this.blobs[i+1:])
			this.blobs = this.blobs[:len(this.blobs)-1]
		} else {
			i++
		}
	}

	return
}

func main() {
	allocator := Constructor(10)
	allocator.Allocate(1, 1)
	allocator.Allocate(1, 2)
	allocator.Allocate(1, 3)
	allocator.Free(2)
	allocator.Allocate(3, 4)
	allocator.Allocate(1, 1)
	allocator.Allocate(1, 1)
	allocator.Free(1)
	fmt.Println(allocator.blobs)
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */
