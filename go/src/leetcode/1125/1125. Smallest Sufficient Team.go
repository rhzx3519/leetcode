package main

import "math/bits"

/*
*
https://leetcode.cn/problems/smallest-sufficient-team/description/

作者：灵茶山艾府
链接：https://leetcode.cn/problems/smallest-sufficient-team/solutions/2214387/zhuang-ya-0-1-bei-bao-cha-biao-fa-vs-shu-qode/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func smallestSufficientTeam(reqSkills []string, people [][]string) (ans []int) {
	m := len(reqSkills)
	sid := make(map[string]int, m)
	for i, s := range reqSkills {
		sid[s] = i // 字符串映射到下标
	}

	n := len(people)
	u, all := 1<<m, 1<<n-1
	f := make([]int, u)
	for j := 1; j < u; j++ {
		f[j] = all // 对应记忆化搜索中的 if (i < 0) return all
	}
	for i, skills := range people {
		mask := 0
		for _, s := range skills { // 把 skills 压缩成一个二进制数 mask[i]
			mask |= 1 << sid[s]
		}
		for j := u - 1; j > 0; j-- {
			r1 := f[j]              // 不选 mask[i]
			r2 := f[j&^mask] | 1<<i // 选 mask[i]
			if bits.OnesCount(uint(r1)) > bits.OnesCount(uint(r2)) {
				f[j] = r2
			}
		}
	}
	res := f[u-1]

	for i := 0; i < n; i++ {
		if res>>i&1 > 0 {
			ans = append(ans, i) // 所有在 res 中的下标
		}
	}
	return
}
