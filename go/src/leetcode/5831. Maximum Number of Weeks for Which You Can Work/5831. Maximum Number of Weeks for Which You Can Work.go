package main

/**
https://leetcode-cn.com/problems/maximum-number-of-weeks-for-which-you-can-work/
https://leetcode-cn.com/problems/maximum-number-of-weeks-for-which-you-can-work/solution/ni-ke-yi-gong-zuo-de-zui-da-zhou-shu-by-rbidw/
 */
func numberOfWeeks(milestones []int) int64 {
	// 耗时最长工作所需周数
	var longest, s int64
	for _, m := range milestones {
		s += int64(m)
		longest = max(longest, int64(m))
	}
	// 其余工作共计所需周数
	var rest = s - longest
	if longest > rest + 1 {
		// 此时无法完成所耗时最长的工作
		return rest * 2 + 1
	} else {
		// 此时可以完成所有工作
		return longest + rest
	}

	var w = int64(0)
	return w
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
