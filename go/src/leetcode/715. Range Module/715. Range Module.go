/**
@author ZhengHao Lou
Date    2021/12/10
*/
package main

/**
https://leetcode-cn.com/problems/range-module/
*/
const MAX_RANGE = 1000000009

type RangeModule struct {
	Root *SegmentNode
}

func Constructor() RangeModule {
	return RangeModule{&SegmentNode{nil, nil, false, false}}
}

func (this *RangeModule) AddRange(left int, right int) {
	this.Root.update(1, MAX_RANGE, left, right-1, true)
}

func (this *RangeModule) QueryRange(left int, right int) bool {
	return this.Root.query(1, MAX_RANGE, left, right-1)
}

func (this *RangeModule) RemoveRange(left int, right int) {
	this.Root.update(1, MAX_RANGE, left, right-1, false)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */

type SegmentNode struct {
	Ls, Rs   *SegmentNode
	Val, Add bool
}

func (node *SegmentNode) update(lc int, rc int, l int, r int, v bool) {
	if l <= lc && rc <= r {
		node.Val, node.Add = v, true
		return
	}
	node.pushdown()
	mid := (lc + rc) >> 1
	if l <= mid {
		node.Ls.update(lc, mid, l, r, v)
	}
	if r > mid {
		node.Rs.update(mid+1, rc, l, r, v)
	}
	node.pushup()
}

func (node *SegmentNode) query(lc int, rc int, l int, r int) bool {
	if l <= lc && rc <= r {
		return node.Val
	}
	node.pushdown()
	mid, ans := (lc+rc)>>1, true
	if l <= mid {
		ans = ans && node.Ls.query(lc, mid, l, r)
	}
	if r > mid {
		ans = ans && node.Rs.query(mid+1, rc, l, r)
	}
	return ans
}

func (node *SegmentNode) pushup() {
	node.Val = node.Ls.Val && node.Rs.Val
}

func (node *SegmentNode) pushdown() {
	if node.Ls == nil {
		node.Ls = &SegmentNode{nil, nil, false, false}
	}
	if node.Rs == nil {
		node.Rs = &SegmentNode{nil, nil, false, false}
	}
	if !node.Add {
		return
	}
	node.Ls.Val, node.Ls.Add, node.Rs.Val, node.Rs.Add = node.Val, true, node.Val, true
	node.Add = false
}

func main() {
	obj := Constructor()
	obj.AddRange(10, 15)
	obj.AddRange(22, 30)
	obj.AddRange(14, 23)
}

/**
 * Your RangeModule object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRange(left,right);
 * param_2 := obj.QueryRange(left,right);
 * obj.RemoveRange(left,right);
 */
