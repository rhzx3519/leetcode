package main

import "fmt"

/**
单调栈
1.首先利用单调栈,计算出每个数作为区间最小值可以往左右两边延拓的长度
2.用上述求出的延拓长度L,去更新答案数组ans[L-1]为其对应最小数字集合(延拓长度为L的数字集合)中最大的一个
3.倒序遍历答案数组,将ans[i]更新max(ans[j]),j>=i

第三步原因解释:
如果存在数字A1,A2,他们的延拓范围分别为L1,L2,且有A1<A2,L1<L2,那么肯定存在长度为L1的子数组其最小值为A2(只需要截取长度为L2且最小值为A2的子数组即可),因此需要倒序完成最终的更新

作者：Gosper
链接：https://leetcode-cn.com/problems/maximum-of-minimum-values-in-all-subarrays/solution/onshi-jian-fu-za-du-dan-diao-zhan-by-gos-uikz/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func findMaximums(nums []int) []int {
	n := len(nums)
	L := make([]int, n)
	for i := 0; i < n; i++ {
		L[i] = n
	}

	var st = []int{}
	for i := range nums {
		for len(st) > 0 && nums[st[len(st)-1]] > nums[i] {
			k := st[len(st)-1]
			L[k] = i // 下标 st[len(st)-1] 右边界为i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	//fmt.Println(L)

	st = []int{}
	for i := n-1; i >= 0; i-- {
		for len(st) > 0 && nums[st[len(st)-1]] > nums[i] {
			k := st[len(st)-1]
			L[k] = L[k]	- i - 1 // 下标 st[len(st)-1] 左边界为i
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}

	//fmt.Println(L)

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[L[i]-1] = max(ans[L[i]-1], nums[i])
	}

	for i := n-2; i >= 0; i-- {
		ans[i]= max(ans[i], ans[i+1])
	}

	fmt.Println(ans)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	findMaximums([]int{0,1,2,4})
	findMaximums([]int{3,4,2,1})
	findMaximums([]int{10,20,50,10})
}