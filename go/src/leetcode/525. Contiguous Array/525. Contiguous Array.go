package main

/**
思路：转化为和=target最长非空子数组长度
 */
func findMaxLength(nums []int) int {
	var ans int
	var s int
	mapper := make(map[int]int)
	mapper[0] = -1

	for i, num := range nums {
		if num == 0 {
			s--
		} else {
			s++
		}

		if last, ok := mapper[s]; !ok {
			mapper[s] = i
		} else {
			ans = max(ans, i-last+1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
