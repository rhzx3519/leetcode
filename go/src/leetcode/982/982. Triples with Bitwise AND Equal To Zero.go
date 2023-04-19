package main

func countTriplets(nums []int) (tot int) {
	cnt := [1 << 16]int{len(nums)} // 统计空集
	for _, x := range nums {
		m := x ^ 0xffff                       // m代表x的补集
		for s := m; s != 0; s = (s - 1) & m { // 枚举m的子集
			cnt[s]++
		}
	}

	for _, x := range nums {
		for _, y := range nums {
			tot += cnt[x&y]
		}
	}
	return
}
