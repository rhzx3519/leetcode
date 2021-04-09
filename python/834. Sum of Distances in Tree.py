class Solution(object):
    def sumOfDistancesInTree(self, N, edges):
        """
        :type N: int
        :type edges: List[List[int]]
        :rtype: List[int]
        """
        adj = [[] for _ in range(N)]
        for a, b in edges:
            adj[a].append(b)
            adj[b].append(a)
        
        distSum = [0]*N # distSum[i]：节点i到它所在子树的节点的距离和，后面更新为：节点i到其他所有节点的距离和
        nodeSum = [1]*N # nodeNum[i]：节点i所在子树的节点个数，保底为1

        def postOrder(root, parent):
            for neighbor in adj[root]:
                if neighbor==parent:
                    continue
                postOrder(neighbor, root)
                nodeSum[root] += nodeSum[neighbor]
                distSum[root] += nodeSum[neighbor] + distSum[neighbor]

        def preOrder(root, parent):
            for neighbor in adj[root]:
                if neighbor==parent: # 如果邻居是自己父亲，跳过。
                    continue
                distSum[neighbor] = distSum[root] - nodeSum[neighbor] + (N - nodeSum[neighbor])
                preOrder(neighbor, root)

        postOrder(0, -1) # dfs的入口。因为N>=1，节点0肯定存在，就从节点0开始搜
        preOrder(0, -1)
        return distSum

if __name__ == '__main__':
    N = 6
    edges = [[0,1],[0,2],[2,3],[2,4],[2,5]]
    su = Solution()
    print su.sumOfDistancesInTree(N, edges)

# https://leetcode-cn.com/problems/sum-of-distances-in-tree/solution/shou-hua-tu-jie-shu-zhong-ju-chi-zhi-he-shu-xing-d/