/**
@author ZhengHao Lou
Date    2023/01/22
*/
package main

/**
https://leetcode.cn/problems/minimum-operations-to-make-array-equal-ii/
*/
func minOperations(nums1 []int, nums2 []int, k int) (tot int64) {
	var x int
	for i := range nums1 {
		if k == 0 {
			if nums1[i] != nums2[i] {
				return -1
			}
			continue
		}
		if (nums2[i]-nums1[i])%k != 0 {
			return -1
		}
		x += (nums2[i] - nums1[i]) / k
		tot += int64(abs((nums2[i] - nums1[i]) / k))
	}
	if x != 0 || tot&1 != 0 {
		return -1
	}
	tot >>= 1
	// fmt.Println(tot)
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	//minOperations([]int{4, 3, 1, 4}, []int{1, 3, 7, 1}, 3)
	//minOperations([]int{3, 8, 5, 2}, []int{2, 4, 1, 6}, 1)
	minOperations([]int{5, 1, 0}, []int{9, 7, 6}, 2)
}
