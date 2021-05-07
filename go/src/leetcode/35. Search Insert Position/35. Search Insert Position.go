package main

import "fmt"

/**
思路：等价于找到第一个大于等于target的元素下标
 */
func searchInsert(nums []int, target int) int {
	n := len(nums)
	ans := n	// ans初始化为n，省略边界判断，有一种情况是target > nums[-1]
	var mid int
	for l, r := 0, n-1; l <= r; {
		mid = (r - l)>>1 + l
		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
	fmt.Println(searchInsert([]int{1,3,5,6}, 0))
	fmt.Println(searchInsert([]int{1,3,3,3,3,4,4,4,5,5,6}, 3))
}