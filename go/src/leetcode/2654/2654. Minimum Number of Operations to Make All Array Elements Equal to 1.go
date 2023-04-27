package main

/*
*
https://leetcode.cn/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
*/
func minOperations(nums []int) (tot int) {
	n := len(nums)
	var c1, gcdAll int
	for _, x := range nums {
		gcdAll = gcd(gcdAll, x)
		if x == 1 {
			c1++
		}
	}
	if gcdAll > 1 {
		return -1
	}
	if c1 > 0 {
		return n - c1
	}

	var minSize = n
	for i := range nums {
		g := 0
		for j, x := range nums[i:] {
			g = gcd(g, x)
			if g == 1 {
				minSize = min(minSize, j+1)
				break
			}
		}
	}
	return (minSize - 1) + (n - 1)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	minOperations([]int{2, 6, 3, 4})
	//minOperations([]int{2, 10, 6, 14})
	//minOperations([]int{6, 10, 15})
}
