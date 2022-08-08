/**
@author ZhengHao Lou
Date    2022/01/04
*/
package main

import "fmt"

/**
https://leetcode-cn.com/problems/maximum-employees-to-be-invited-to-a-meeting/
https://leetcode-cn.com/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/
思路：基环树(pseudotree)
time: O(n)
space: O(n)

下面介绍基环树问题的通用写法。

我们可以通过一次拓扑排序「剪掉」所有树枝，因为拓扑排序后，树枝节点的入度均为 00，基环节点的入度均为 11。这样就可以将基环和树枝分开，从而简化后续处理流程：

如果要遍历基环，可以从拓扑排序后入度为 11 的节点出发，在图上搜索；
如果要遍历树枝，可以以基环与树枝的连接处为起点，顺着反图来搜索树枝（搜索入度为 00 的节点），从而将问题转化成一个树形问题。
对于本题，我们可以遍历所有基环，并按基环大小分类计算：

对于大小大于 22 的基环，我们取基环大小的最大值；
对于大小等于 22 的基环，我们可以从基环上的点出发，在反图上找到最大的树枝节点深度。
时间复杂度和空间复杂度均为 O(n)O(n)。

作者：endlesscheng
链接：https://leetcode-cn.com/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
 */
func maximumInvitations(favor []int) int {
	n := len(favor)
	ind := make([]int, n)
	depth := make([]int, n)	// depth[i]表示以i为根结点的favor链的最大长度
	for _, i := range favor {
		ind[i]++
	}

	que := []int{}
	for i := range ind {	// 入度为0的节点入队
		if ind[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) != 0 {	// 拓扑排序，减去环上的所有树枝
		r := que[0]
		que = que[1:]
		depth[r]++
		w := favor[r]
		depth[w] = max(depth[w], depth[r])
		if ind[w]--; ind[w] == 0 {
			que = append(que, w)
		}
	}

	var maxRingSize, sumChainSize int

	for i := range ind {
		if ind[i] == 0 {
			continue
		}
		// i为基环上的点，拓扑排序之后入度>0
		ind[i] = 0
		ringSize := 1
		for v := favor[i]; v != i; v = favor[v] {
			ind[v] = 0 // 入度标记为0，避免重复访问
			ringSize++
		}

		if ringSize == 2 {	// 二元环，答案为两个节点的最大深度之和
			sumChainSize += depth[i] + depth[favor[i]] + 2 // 累加两条链的长度(如果图中存在 多个 这样的情况，我们可以把他们 都安排上)
		} else {	// 基环节点数量大于2, 基环的大小就是答案
			maxRingSize = max(maxRingSize,  ringSize)
		}
	}

	return max(maxRingSize, sumChainSize)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumInvitations([]int{2,2,1,2}))
	fmt.Println(maximumInvitations([]int{1,2,0}))
	fmt.Println(maximumInvitations([]int{3,0,1,4,1}))
}

