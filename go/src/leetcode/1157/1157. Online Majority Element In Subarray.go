package main

import (
	"fmt"
	"sort"
)

type MajorityChecker struct {
	root seg
	cnt  map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	root := newSegmentTree(arr)
	cnt := make(map[int][]int)
	for i, x := range arr {
		cnt[x] = append(cnt[x], i)
	}

	return MajorityChecker{
		root: root,
		cnt:  cnt,
	}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) (tot int) {

	tot = -1
	m := this.root.query(1, left+1, right+1)
	l := sort.Search(len(this.cnt[m.x]), func(i int) bool {
		return this.cnt[m.x][i] >= left
	})
	r := sort.Search(len(this.cnt[m.x]), func(i int) bool {
		return this.cnt[m.x][i] > right
	}) - 1
	defer func() {
		fmt.Println(tot, m.x, this.cnt[m.x], l, r)
	}()

	if r-l+1 >= threshold {
		tot = m.x
	}
	return
}

func main() {
	c := Constructor([]int{2, 1, 1, 1, 2, 1, 2, 1, 2, 2, 1, 1, 2})
	c.Query(9, 11, 2) // 2,2,1
	//c.Query(0, 3, 3)
	//c.Query(2, 3, 2)
}

type pair struct {
	x, c int
}

// l 和 r 也可以写到方法参数上，实测二者在执行效率上无异
// 考虑到 debug 和 bug free 上的优点，写到结构体参数中
type seg []struct {
	l, r int
	mode pair
}

// 单点更新：build 和 update 通用
func (t seg) set(o, val int) {
	t[o].mode = pair{val, 1}
}

// 合并两个节点上的数据：maintain 和 query 通用
// 要求操作满足区间可加性
// 例如 + * | & ^ min max gcd mulMatrix 摩尔投票 最大子段和 ...
func (seg) op(a, b pair) pair {
	if a.x == b.x {
		return pair{a.x, a.c + b.c}
	}

	if a.c > b.c {
		return pair{a.x, a.c - b.c}
	}
	return pair{b.x, b.c - a.c}
}

func (t seg) maintain(o int) {
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].mode = t.op(lo.mode, ro.mode)
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t.set(o, a[l-1])
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

// o=1  1<=i<=n
func (t seg) update(o, i, val int) {
	if t[o].l == t[o].r {
		t.set(o, val)
		return
	}
	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i, val)
	} else {
		t.update(o<<1|1, i, val)
	}
	t.maintain(o)
}

// o=1  [l,r] 1<=l<=r<=n
func (t seg) query(o, l, r int) pair {
	if l <= t[o].l && t[o].r <= r {
		return t[o].mode
	}
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	vl := t.query(o<<1, l, r)
	vr := t.query(o<<1|1, l, r)
	return t.op(vl, vr)
}

func (t seg) queryAll() int { return t[1].mode.x }

// a 不能为空
func newSegmentTree(a []int) seg {
	t := make(seg, 4*len(a))
	t.build(a, 1, 1, len(a))
	return t
}
