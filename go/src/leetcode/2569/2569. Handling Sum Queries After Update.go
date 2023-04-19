package main

import "fmt"

/*
*
https://leetcode.cn/problems/handling-sum-queries-after-update/
(x+p) + p + p
*/
type seg []struct {
	l, r, cnt1 int
	flip       bool
}

func (t seg) maintain(o int) { t[o].cnt1 = t[o<<1].cnt1 + t[o<<1|1].cnt1 }

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].cnt1 = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	t.maintain(o)
}

func (t seg) do(O int) {
	o := &t[O]
	o.cnt1 = o.r - o.l + 1 - o.cnt1
	o.flip = !o.flip
}

func (t seg) spread(o int) {
	if t[o].flip {
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].flip = false
	}
}

func (t seg) update(o, l, r int) {
	if l <= t[o].l && t[o].r <= r {
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if l <= m {
		t.update(o<<1, l, r)
	}
	if m < r {
		t.update(o<<1|1, l, r)
	}
	t.maintain(o)
}

func handleQuery(nums1, nums2 []int, queries [][]int) (ans []int64) {
	sum := 0
	for _, x := range nums2 {
		sum += x
	}
	t := make(seg, len(nums1)*4)
	t.build(nums1, 1, 1, len(nums1))
	for _, q := range queries {
		if q[0] == 1 {
			t.update(1, q[1]+1, q[2]+1)
		} else if q[0] == 2 {
			sum += q[1] * t[1].cnt1
		} else {
			ans = append(ans, int64(sum))
		}
	}
	return
}

func main() {
	fmt.Println(handleQuery([]int{1, 0, 1}, []int{0, 0, 0}, [][]int{{1, 1, 1}, {2, 1, 0}, {3, 0, 0}}))
	fmt.Println(handleQuery([]int{1}, []int{5}, [][]int{{2, 0, 0}, {3, 0, 0}}))
}
