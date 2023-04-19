package main

/*
*
https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/
*/
func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right] >= arr[right-1] {
		right--
	}
	if right == 0 {
		return 0
	}
	var ans = n
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[left] > arr[right] {
			right++
		}
		ans = min(ans, right-left-1) // 删除arr[left+1:right]
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})
}
