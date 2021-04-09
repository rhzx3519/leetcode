class TreeAncestor:

    def __init__(self, n, parent):
        self.cols = 20      # log(50000) < 20

        self.dp = [[-1] * self.cols for _ in range(n)]
        for i in range(n):
            self.dp[i][0] = parent[i]
        # 动态规划设置祖先, dp[node][j] 表示 node 往前推第 2^j 个祖先
        for j in range(1, self.cols):
            for i in range(n):
                if self.dp[i][j-1] != -1:
                    self.dp[i][j] = self.dp[self.dp[i][j-1]][j-1]

    def getKthAncestor(self, node, k):
        for i in range(self.cols - 1, -1, -1):
            if k & (1 << i):
                node = self.dp[node][i]
                if node == -1:
                    break
        return node

# Your TreeAncestor object will be instantiated and called as such:
# obj = TreeAncestor(n, parent)
# param_1 = obj.getKthAncestor(node,k)

'''
解题思路
一、DP倍增法，本质类似于二分法（和这个题有异曲同工之妙: 求一个数的n次方）

1. 先看dp递推公式吧，看了可能会有点懵，不过动态规划就是这样的，不容易想到。

状态定义: dp[node][j]，表示 node 的第 2^j 个祖先。其中 node取值为 0~n-1, 列取值为 0~20
# 题目限定范围 log(50000) < 20
递推公式: dp[ node ][j] = dp[ dp[node][j-1] ][j-1]

2. 解题步骤如下:

a. 首先根据题目条件和递推公式，构造二维DP数组
b. 求给定 node 的第 k 个祖先

3. 举例说明

当k=10 时，二进制为 1010, 等于 2^3 + 2^1
当求node的第10个祖先时，结果为: dp[dp[node][2^3]][2^1]，和幂运算类似，底数不变，指数相加

总结：据说这个题是典型的ACM题，如果不清楚这个算法，是很难做出来的，我也是参考了别人的题解，汗。。。

作者：dz-lee
链接：https://leetcode-cn.com/problems/kth-ancestor-of-a-tree-node/solution/dpbei-zeng-fa-dai-ma-qing-xi-python3xiang-xi-zhu-s/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
'''