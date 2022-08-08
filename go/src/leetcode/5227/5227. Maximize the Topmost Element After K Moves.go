/**
@author ZhengHao Lou
Date    2022/03/14
*/
package main

/**
https://leetcode-cn.com/problems/maximize-the-topmost-element-after-k-moves/
设 \textit{nums}nums 的长度为 nn，根据题意：

如果 n=1n=1，那么我们只能在栈不为空时删除栈顶，栈为空时将 \textit{nums}[0]nums[0] 入栈。因此 kk 为奇数时，kk 次操作后栈为空，返回 -1−1；kk 为偶数时则返回 \textit{nums}[0]nums[0]。
如果 k=0k=0，无法执行任何操作，直接返回 \textit{nums}[0]nums[0]。
其余情况，按数组元素的下标 ii 分类讨论：

如果 i=0i=0，我们可以不断地删除-添加 \textit{nums}[0]nums[0]，如果 kk 为偶数，那么最后栈顶为 \textit{nums}[0]nums[0]；如果 kk 为奇数（这里要求 k>1k>1），我们可以在倒数第二步删除 \textit{nums}[1]nums[1]，最后一步将 \textit{nums}[0]nums[0] 入栈，从而保证 \textit{nums}[0]nums[0] 可以为栈顶。
如果 0<i<k-10<i<k−1，我们仍然可以仿造上述流程操作。
如果 i=k-1i=k−1，最后一步操作只能删除 \textit{nums}[i]nums[i]，所以无法将 \textit{nums}[i]nums[i] 置于栈顶。
如果 i=ki=k，那么可以删除前 kk 个元素，将 \textit{nums}[i]nums[i] 置于栈顶。
如果 i>ki>k，\textit{nums}[i]nums[i] 前面的元素无法删除，所以无法将 \textit{nums}[i]nums[i] 置于栈顶。
综上所述，我们可以让 i<k-1i<k−1 或 i=ki=k 的数组元素作为 kk 次操作后的栈顶。这些元素的最大值即为答案。

作者：endlesscheng
链接：https://leetcode-cn.com/problems/maximize-the-topmost-element-after-k-moves/solution/by-endlesscheng-vmtr/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func maximumTop(nums []int, k int) int {
	if k == 0 {
		return nums[0]
	}
	n := len(nums)
	if n == 1 {
		if k&1 == 0 {
			return nums[0]
		}
		return -1
	}

	var mx = -1
	for i := 0; i < min(n, k-1); i++ {
		mx = max(mx, nums[i])
	}
	if k < n {
		mx = max(mx, nums[k])
	}
	return mx
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
