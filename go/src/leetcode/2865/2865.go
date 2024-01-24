package main

/*
*
https://leetcode.cn/problems/beautiful-towers-i/?envType=daily-question&envId=2024-01-24
*/
func maximumSumOfHeights(maxHeights []int) (tot int64) {
	n := len(maxHeights)
	for i := range maxHeights {
		x := maxHeights[i]
		var t = x
		for j := i - 1; j >= 0; j-- {
			x = min(x, maxHeights[j])
			t += x
		}
		x = maxHeights[i]
		for k := i + 1; k < n; k++ {
			x = min(x, maxHeights[k])
			t += x
		}
		tot = max(tot, int64(t))
	}
	return
}
