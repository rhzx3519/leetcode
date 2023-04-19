package main

import (
	"fmt"
	"sort"
)

/*
*
https://leetcode.cn/problems/minimum-time-to-complete-all-tasks/
思路：贪心
1. 按照结束时间排序
2. 如果区间内电脑已在运行，优先使用这些时间点；否则安排在后缀运行，新增时间点
*/
func findMinimumTime(tasks [][]int) (tot int) {
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][1] != tasks[j][1] {
			return tasks[i][1] < tasks[j][1]
		}
		return tasks[i][0] < tasks[j][0]
	})

	fmt.Println(tasks)
	run := make([]bool, tasks[len(tasks)-1][1]+1)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		for _, b := range run[start : end+1] {
			if b {
				d--
			}
			if d == 0 {
				break
			}
		}
		for i := end; d > 0; i-- {
			if !run[i] {
				run[i] = true
				d--
				tot++
			}
		}
	}

	return
}

/*
*
线段树优化
*/
func findMinimumTime1(tasks [][]int) (tot int) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	u := tasks[len(tasks)-1][1]
	st := make(seg, 4*u)
	st.build(1, 1, u)
	for _, t := range tasks {
		start, end, d := t[0], t[1], t[2]
		d -= st.query(1, start, end)
		if d > 0 {
			st.update(1, start, end, &d)
		}
	}
	return st[1].cnt
}

type seg []struct {
	l, r, cnt int
	todo      bool
}

func (t seg) build(o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		return
	}
	m := (l + r) >> 1
	t.build(o<<1, l, m)
	t.build(o<<1|1, m+1, r)
}

func (t seg) do(i int) {
	o := &t[i]
	o.cnt = o.r - o.l + 1
	o.todo = true
}

func (t seg) spread(o int) {
	if t[o].todo {
		t[o].todo = false
		t.do(o << 1)
		t.do(o<<1 | 1)
	}
}

// 查询区间 [l,r]   o=1
func (t seg) query(o, l, r int) int {
	if l <= t[o].l && t[o].r <= r {
		return t[o].cnt
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r <= m {
		return t.query(o<<1, l, r)
	}
	if m < l {
		return t.query(o<<1|1, l, r)
	}
	return t.query(o<<1, l, r) + t.query(o<<1|1, l, r)
}

func (t seg) update(o, l, r int, suffix *int) {
	size := t[o].r - t[o].l + 1
	if t[o].cnt == size {
		return
	}
	if l <= t[o].l && t[o].r <= r && size-t[o].cnt <= *suffix {
		*suffix -= size - t[o].cnt
		t.do(o)
		return
	}
	t.spread(o)
	m := (t[o].l + t[o].r) >> 1
	if r > m { // 优先更新右子树
		t.update(o<<1|1, l, r, suffix)
	}
	if *suffix > 0 {
		t.update(o<<1, l, r, suffix)
	}
	t[o].cnt = t[o<<1].cnt + t[o<<1|1].cnt
}

func main() {
	//findMinimumTime([][]int{{2, 3, 1}, {4, 5, 1}, {1, 5, 2}})
	//findMinimumTime([][]int{{1, 3, 2}, {2, 5, 3}, {5, 6, 2}})
	//findMinimumTime([][]int{{1, 1, 1}})
	findMinimumTime1([][]int{{8, 19, 1}, {3, 20, 1}, {1, 20, 2}, {6, 13, 3}})
}
