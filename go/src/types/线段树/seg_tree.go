package 线段树

/**
lazy 线段树

问题：一个数组，更新一个子数组（都加上一个数、把子数组内的元素取反...）
			 查询一个子数组 (查询、求最大值)

两大思想:
1. 挑选O(n)个特殊区间, 使得任意一个区间可以拆分成O(logn)个特殊区间（使用最近公共祖先来思考）
	O(n) <= 4n
挑选O(n)个特殊区间: build

2. lazy更新 / 延迟更新
lazy tag: 用一个数组维护每个区间需要更新的值
如果这个值 = 0, 表示不需要更新
如果这个值 != 0, 表示更新操作在这个区间停住了，不继续递归更新子数组

如果后来又来了一个更新操作，破坏了lazy tag的区间，那么这个区间就得继续更新了

// 讲解视频
https://www.bilibili.com/video/BV15D4y1G7ms/?vd_source=330d34de7eb65adb5cad7c60bd5b78ec
// 模版代码
https://github.com/EndlessCheng/codeforces-go/blob/ae6b5202069774f93266b24324a085e6e607db90/copypasta/segment_tree.go#L125

*/

type seg []struct {
	l, r, cnt1 int
	flip       bool // lazy tag
}

func (t seg) build(a []int, o, l, r int) {
	t[o].l, t[o].r = l, r
	if l == r {
		t[o].cnt1 = a[l-1]
		return
	}
	m := (l + r) >> 1
	t.build(a, o<<1, l, m)
	t.build(a, o<<1|1, m+1, r)
	// 维护 ...
	t.maintain(o)
}

func (t seg) do(O int) { // 更新本节点的lazy tag
	o := &t[O]
	o.cnt1 = o.r - o.l + 1 - o.cnt1
	o.flip = !o.flip
}

func (t seg) spread(o int) {
	if t[o].flip { // 破坏了原有的lazy tag，继续递归更新子树
		t.do(o << 1)
		t.do(o<<1 | 1)
		t[o].flip = false
	}
}

func (t seg) maintain(o int) { //
	t[o].cnt1 = t[o<<1].cnt1 + t[o<<1|1].cnt1
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
	var sum int
	for i := range nums2 {
		sum += nums2[i]
	}
	n := len(nums1)
	t := make(seg, n*4)
	t.build(nums1, 1, 1, n)
	for _, q := range queries {
		if q[0] == 1 {
			t.update(1, q[1]+1, q[2]+1)
		} else if q[0] == 2 {
			sum += t[1].cnt1 * q[1]
		} else {
			ans = append(ans, int64(sum))
		}
	}
	return
}
