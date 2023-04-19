package main

import "fmt"

/**
1. 首先从后向前查找第一个顺序对 (i,i+1)，满足 a[i] < a[i+1]。
这样「较小数」即为 a[i]。此时 [i+1,n) 必然是下降序列。

2. 如果找到了顺序对，那么在区间 [i+1,n) 中从后向前查找第一个元素 j 满足 a[i] < a[j]。
这样「较大数」即为 a[j]。

3. 交换 a[i]与 a[j]，此时可以证明区间 [i+1,n) 必为降序。我们可以直接使用双指针反转区间 [i+1,n)
使其变为升序，而无需对该区间进行排序。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func nextPermutation(nums []int)  {
	n := len(nums)
	i := n-2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n-1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l, r := i+1, n-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}

func main() {
	nums := []int{1,2,3}
	for i := 0; i < 6; i++ {
		nextPermutation(nums)
		fmt.Println(nums)
	}
}
