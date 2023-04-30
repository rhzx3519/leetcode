package main

// 树状数组模板
type BIT []int

// 将下标 i 上的数加一
func (t BIT) inc(i int) {
	for ; i < len(t); i += i & -i {
		t[i]++
	}
}

// 返回闭区间 [1, i] 的元素和
func (t BIT) sum(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += t[i]
	}
	return
}

// 返回闭区间 [left, right] 的元素和
func (t BIT) query(left, right int) int {
	return t.sum(right) - t.sum(left-1)
}
