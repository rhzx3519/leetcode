package main

var s []byte

type seg []struct {
	l, r, pre, suf, max int
}

func (t seg) maintain(o int) { //
	lo, ro := t[o<<1], t[o<<1|1]
	t[o].pre = lo.pre
	t[o].suf = ro.suf
	t[o].max = max(lo.max, ro.max)
	if s[lo.r-1] == s[lo.r] {
		if lo.suf == lo.r-lo.l+1 {
			t[o].pre += ro.pre
		}
		if ro.pre == ro.r-ro.l+1 {
			t[o].suf += lo.suf
		}
		t[o].max = max(t[o].max, lo.suf+ro.pre)
	}
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].pre = 1
		t[o].suf = 1
		t[o].max = 1
		return
	}

	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
	// 维护 ...
	t.maintain(o)
}

func (t seg) update(o, i int) {
	if t[o].l == t[o].r {
		return
	}

	m := (t[o].l + t[o].r) >> 1
	if i <= m {
		t.update(o<<1, i)
	} else {
		t.update(o<<1|1, i)
	}
	t.maintain(o)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestRepeating(S string, queryCharacters string, queryIndices []int) []int {
	s = []byte(S)
	n := len(s)
	t := make(seg, n*4)
	t.build(1, 1, n)
	ans := make([]int, len(queryIndices))
	for i, index := range queryIndices {
		s[index] = queryCharacters[i]
		t.update(1, index+1)
		ans[i] = t[1].max
	}
	return ans
}
